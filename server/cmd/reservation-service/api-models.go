package main

import (
	"time"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// createReservationRequest represents deserialized data from CreateReservation request
type createReservationRequest struct {
	FromTime   time.Time `json:"from"`
	ToTime     time.Time `json:"to"`
	ChildCount int       `json:"numberOfChildren"`
	AdultCount int       `json:"numberOfAdults"`
}

// toReservation maps a request to add an Reservation to model Reservation
func (c createReservationRequest) toReservation() *bookly.Reservation {
	return &bookly.Reservation{
		FromTime:   c.FromTime,
		ToTime:     c.ToTime,
		AdultCount: c.AdultCount,
		ChildCount: c.ChildCount,
	}
}
