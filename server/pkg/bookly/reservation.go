package bookly

import (
	"context"
	"time"
)

// Reservation holds information about client reservation
type Reservation struct {
	ID         int64
	ClientID   int64
	HotelID    int64
	OfferID    int64
	FromTime   time.Time
	ToTime     time.Time
	ChildCount int
	AdultCount int
}

// ReservationObject holds data about reservations for client
type ReservationObject struct {
	HotelPreview HotelInfoPreview `json:"hotelInfoPreview"`
	Reservation  ReservationInfo  `json:"reservationInfo"`
	OfferPreview OfferInfoPreview `json:"offerInfoPreview"`
}

// ReservationInfo holds info about offer for reservation purpose
type ReservationInfo struct {
	ReservationID int64     `json:"reservationID"`
	FromTime      time.Time `json:"from"`
	ToTime        time.Time `json:"to"`
	ChildAmount   int       `json:"numberOfChildren"`
	AdultAmount   int       `json:"numberOfAdults"`
	ReviewID      *int64    `json:"reviewID,omitempty"`
}

// ReservationStorage is responsible for storage operations on client reservation
type ReservationStorage interface {
	CreateReservation(ctx context.Context, reservation *Reservation) (int64, error)
	DeleteReservation(ctx context.Context, reservationID int64) error
	GetSpecificReservation(ctx context.Context, reservationID int64) (*Reservation, error)
	GetClientReservations(ctx context.Context, clientID int64) ([]*Reservation, error)
	GetHotelReservations(ctx context.Context, hotelID int64) ([]*Reservation, error)
	IsReservationOwnedByClient(ctx context.Context, clientID int64, reservationID int64) (bool, error)
}

// ReservationService is a service which is responsible for actions related to reservations
type ReservationService interface {
	CreateReservation(ctx context.Context, reservation *Reservation) error
	GetClientReservations(ctx context.Context, clientID int64, pageNumber int, pageSize int) ([]*ReservationObject, error)
	DeleteReservation(ctx context.Context, clientID int64, reservationID int64) error
}
