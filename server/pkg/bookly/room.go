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
	GetRoom(ctx context.Context, roomNumber string, hotelID int64) (Room, error)
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
