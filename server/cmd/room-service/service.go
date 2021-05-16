package main

import (
	"context"
	"fmt"

	"github.com/pw-software-engineering/b-team/server/pkg/paging"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type roomService struct {
	roomStorage  bookly.RoomStorage
	offerStorage bookly.OfferStorage
}

func newRoomService(roomStorage bookly.RoomStorage, offerStorage bookly.OfferStorage) *roomService {
	return &roomService{
		roomStorage:  roomStorage,
		offerStorage: offerStorage,
	}
}

// CreateRoom validates requested room and adds it storage
func (r *roomService) CreateRoom(ctx context.Context, room bookly.Room, hotelID int64) (int64, error) {
	err := validateRoom(room)
	if err != nil {
		return 0, err
	}

	return r.roomStorage.CreateRoom(ctx, room, hotelID)
}

// DeleteRoom transfer task to room storage
func (r *roomService) DeleteRoom(ctx context.Context, roomID int64, hotelID int64) error {
	return r.roomStorage.DeleteRoom(ctx, roomID, hotelID)
}

// GetAllHotelRooms implements split rooms into pages and business logic of API params
func (r *roomService) GetAllHotelRooms(ctx context.Context, hotelID int64, pageNumber int, pageSize int, filter string) ([]*bookly.Room, error) {
	if filter != "" {
		room, err := r.roomStorage.GetRoomByName(ctx, filter, hotelID)
		if err != nil {
			return nil, err
		}
		tab := []*bookly.Room{}
		tab = append(tab, &room)
		return tab, nil
	}

	if pageNumber < 0 || pageSize < 0 {
		return nil, fmt.Errorf("Room service: Wrong pages parametrs")
	} else if pageNumber > 0 && pageSize > 0 {
		result, err := r.roomStorage.GetAllHotelRooms(ctx, hotelID)
		if err != nil {
			return nil, err
		}

		start, end := paging.GetPageItems(pageNumber, pageSize, len(result))
		result = result[start:end]

		return result, nil
	}

	result, err := r.roomStorage.GetAllHotelRooms(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRoomsRelatedWithOffer implements business logic of feat getting rooms related with offer
func (r *roomService) GetRoomsRelatedWithOffer(ctx context.Context, offerID int64, hotelID int64, pageNumber int, pageSize int, filter string) ([]*bookly.Room, error) {
	check, err := r.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if err != nil {
		return nil, err
	}
	if check == false {
		return nil, bookly.ErrOfferNotOwned
	}

	roomsID, err := r.roomStorage.GetRoomsRelatedWithOffer(ctx, offerID)
	if err != nil {
		return nil, err
	}

	var rooms []*bookly.Room
	for _, v := range roomsID {
		room, err := r.roomStorage.GetRoom(ctx, v)
		if err != nil {
			return nil, err
		}
		room.OfferID, err = r.roomStorage.GetOffersRelatedWithRoom(ctx, v)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}

	if filter != "" {
		index := -1
		for i, v := range rooms {
			if filter == v.RoomNumber {
				index = i
				break
			}
		}
		if index == -1 {
			return nil, bookly.ErrRoomNotFound
		}
		return []*bookly.Room{rooms[index]}, nil
	}

	if pageNumber < 0 || pageSize < 0 {
		return nil, fmt.Errorf("Room service: Wrong pages parametrs")
	} else if pageNumber > 0 && pageSize > 0 {
		start, end := paging.GetPageItems(pageNumber, pageSize, len(rooms))
		rooms = rooms[start:end]
	}

	return rooms, nil
}

// AddRoomToOffer implements business logic of feat adding rooms to offer
func (r *roomService) AddRoomToOffer(ctx context.Context, offerID int64, roomID int64, hotelID int64) error {
	check, err := r.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if err != nil {
		return err
	}
	if check == false {
		return bookly.ErrOfferNotOwned
	}

	check, err = r.roomStorage.IsRoomOwnedByHotel(ctx, roomID, hotelID)
	if err != nil {
		return err
	}
	if check == false {
		return bookly.ErrRoomNotOwnedByHotel
	}

	check, err = r.roomStorage.IsExistLinkWithRoomAndOffer(ctx, offerID, roomID)
	if err != nil {
		return err
	}
	if check == false {
		err = r.roomStorage.AddLinkWithRoomAndOffer(ctx, offerID, roomID)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteRoomFromOffer implements business logic of feat delete rooms from offer
func (r *roomService) DeleteRoomFromOffer(ctx context.Context, offerID int64, roomID int64, hotelID int64) error {
	check, err := r.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if err != nil {
		return err
	}
	if check == false {
		return bookly.ErrOfferNotOwned
	}

	check, err = r.roomStorage.IsRoomOwnedByHotel(ctx, roomID, hotelID)
	if err != nil {
		return err
	}
	if check == false {
		return bookly.ErrRoomNotOwnedByHotel
	}

	err = r.roomStorage.DeleteLinkWithRoomAndOffer(ctx, offerID, roomID)
	if err != nil {
		return err
	}

	return nil
}

var _ bookly.RoomService = &roomService{}

func validateRoom(room bookly.Room) error {
	if len(room.RoomNumber) < 1 || len(room.RoomNumber) > 254 {
		return fmt.Errorf("RoomNumber: Invalidate, wrong size")
	}
	return nil
}
