package bookly

import (
	"context"
)

// Room is a model for room, not implemented yed
type Room struct {
	ID         int64
	RoomNumber string
	HotelID    int64
}

// RoomStorage is responsible for persisting offers
type RoomStorage interface {
	CreateRoom(ctx context.Context, room Room, hotelID int64) (int64, error)
	DeleteRoom(ctx context.Context, roomID int64, hotelID int64) (DeleteResponse, error)
	GetAllHotelRooms(ctx context.Context, hotelID int) ([]*Room, error)
}

// RoomService is a service which is responsible for actions related to offers
type RoomService interface{}

// DeleteResponse enum type of respond
type DeleteResponse int

// Enum of DeleteResponse
const (
	RoomNotFound DeleteResponse = iota
	RoomNotBelongToHotel
	RoomSuccess
	RoomError
)
