package main

import (
	"context"
	"fmt"

	"github.com/pw-software-engineering/b-team/server/pkg/paging"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type roomService struct {
	roomStorage bookly.RoomStorage
}

func newRoomService(roomStorage bookly.RoomStorage) *roomService {
	return &roomService{roomStorage: roomStorage}
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
		room, err := r.roomStorage.GetRoom(ctx, filter, hotelID)
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

var _ bookly.RoomService = &roomService{}

func validateRoom(room bookly.Room) error {
	if len(room.RoomNumber) < 1 || len(room.RoomNumber) > 254 {
		return fmt.Errorf("RoomNumber: Invalidate, wrong size")
	}
	return nil
}
