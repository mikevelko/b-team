package bookly

import (
	"context"
	"errors"
)

// Room is a model for room, not implemented yed
type Room struct {
	ID         int64   `json:"roomID"`
	RoomNumber string  `json:"hotelRoomNumber"`
	HotelID    int64   `json:"-"`
	OfferID    []int64 `json:"offerID"`
}

// RoomStorage is responsible for persisting offers
type RoomStorage interface {
	CreateRoom(ctx context.Context, room Room, hotelID int64) (int64, error)
	DeleteRoom(ctx context.Context, roomID int64, hotelID int64) error
	GetAllHotelRooms(ctx context.Context, hotelID int64) ([]*Room, error)
	GetRoomByName(ctx context.Context, roomNumber string, hotelID int64) (Room, error)

	GetRoom(ctx context.Context, roomID int64) (Room, error)
	OffersRelatedWithRoom(ctx context.Context, roomID int64) ([]int64, error)
	RoomsRelatedWithRoom(ctx context.Context, offerID int64) ([]int64, error)
	AddLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) error
	DeleteLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) error
	IsExistLinkWithRoomAndOffer(ctx context.Context, offerID int64, roomID int64) (bool, error)
}

// RoomService is a service which is responsible for actions related to offers
type RoomService interface {
	CreateRoom(ctx context.Context, room Room, hotelID int64) (int64, error)
	DeleteRoom(ctx context.Context, roomID int64, hotelID int64) error
	GetAllHotelRooms(ctx context.Context, hotelID int64, pageNumber int, pageSize int, filter string) ([]*Room, error)
}

// ErrRoomNotFound indicates that room doesn't exist in database
var ErrRoomNotFound = errors.New("room not fount")

// ErrRoomNotBelongToHotel indicates that room not belong to this hotel
var ErrRoomNotBelongToHotel = errors.New("room not belong to hotel")

// ErrRoomAlreadyExists indicates that room already exist in database
var ErrRoomAlreadyExists = errors.New("room already exists")

// ErrRoomIsRelatedWithOffer indicates that room is related with offer
var ErrRoomIsRelatedWithOffer = errors.New("room is related with offer")

// ErrLinkOfferRoomNotFound indicates that record with offer and room link not found
var ErrLinkOfferRoomNotFound = errors.New("record with offer and room link not found")
