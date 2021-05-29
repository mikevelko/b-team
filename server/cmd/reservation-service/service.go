package main

import (
	"context"
	"time"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/paging"
)

type reservationService struct {
	offerStorage       bookly.OfferStorage
	reservationStorage bookly.ReservationStorage
	hotelStorage       bookly.HotelStorage
}

func newReservationService(offers bookly.OfferStorage,
	reservations bookly.ReservationStorage,
	hotels bookly.HotelStorage) *reservationService {
	return &reservationService{reservationStorage: reservations, offerStorage: offers, hotelStorage: hotels}
}

// CreateReservation handles business logic connected to creating reservations
func (s *reservationService) CreateReservation(ctx context.Context, reservation *bookly.Reservation) error {
	if reservation.ToTime.Before(reservation.FromTime) {
		// Offers is not available during negative time periods
		return bookly.ErrOfferNotAvailable
	}
	isActive, err := s.offerStorage.IsOfferActive(ctx, reservation.OfferID)
	if err != nil {
		return bookly.ErrOfferNotFound
	}
	if !isActive {
		return bookly.ErrOfferNotAvailable
	}
	offer, errGet := s.offerStorage.GetSpecificOffer(ctx, reservation.OfferID)
	if errGet != nil {
		return bookly.ErrOfferNotFound
	}
	if offer.MaxGuests < reservation.ChildCount+reservation.AdultCount {
		// offer won't be able to book that amount of people
		return bookly.ErrReservationTooBig
	}
	// todo: check availability intervals
	_, errCreate := s.reservationStorage.CreateReservation(ctx, reservation)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// GetClientReservations retrieves client reservations
func (s *reservationService) GetClientReservations(ctx context.Context, clientID int64, pageNumber int, pageSize int) ([]*bookly.ReservationObject, error) {
	reservations, errGet := s.reservationStorage.GetClientReservations(ctx, clientID)
	if errGet != nil {
		return nil, errGet
	}

	list := []*bookly.ReservationObject{}
	for _, el := range reservations {
		obj := &bookly.ReservationObject{}
		offer, err := s.offerStorage.GetSpecificOffer(ctx, el.OfferID)
		if err != nil {
			return nil, err
		}
		hotel, errHotel := s.hotelStorage.GetHotelDetails(ctx, el.HotelID)
		if errHotel != nil {
			return nil, errHotel
		}
		obj.OfferPreview.OfferID = el.OfferID
		obj.OfferPreview.OfferTitle = offer.OfferTitle
		obj.OfferPreview.OfferPreviewPicture = offer.OfferPreviewPicture
		// todo: add reviews
		obj.Reservation.ReviewID = nil
		obj.Reservation.ReservationID = el.ID
		obj.Reservation.FromTime = el.FromTime
		obj.Reservation.ToTime = el.ToTime
		obj.Reservation.AdultAmount = el.AdultCount
		obj.Reservation.ChildAmount = el.ChildCount

		obj.HotelPreview.HotelID = el.HotelID
		obj.HotelPreview.HotelName = hotel.Name
		obj.HotelPreview.Country = hotel.Country
		obj.HotelPreview.City = hotel.City
		list = append(list, obj)
	}
	start, end := paging.GetPageItems(pageNumber, pageSize, len(list))
	return list[start:end], nil
}

// DeleteReservation deletes reservation provided that it can be deleted
func (s *reservationService) DeleteReservation(ctx context.Context, clientID int64, reservationID int64) error {
	isOwned, err := s.reservationStorage.IsReservationOwnedByClient(ctx, clientID, reservationID)
	if err != nil {
		return bookly.ErrReservationDoesNotExists
	}
	if !isOwned {
		return bookly.ErrReservationNotOwned
	}
	// todo: add hotel and offer verification (if actually this needs to be checked)
	reservation, errGet := s.reservationStorage.GetSpecificReservation(ctx, reservationID)
	if errGet != nil {
		return bookly.ErrReservationDoesNotExists
	}
	if time.Now().After(reservation.FromTime) {
		return bookly.ErrReservationInProgress
	}
	errDelete := s.reservationStorage.DeleteReservation(ctx, reservationID)
	return errDelete
}
