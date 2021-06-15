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
	reviewStorage      bookly.ReviewStorage
	roomStorage        bookly.RoomStorage
	userStorage        bookly.UserStorage
}

func newReservationService(offers bookly.OfferStorage,
	reservations bookly.ReservationStorage,
	hotels bookly.HotelStorage,
	reviews bookly.ReviewStorage,
	rooms bookly.RoomStorage,
	users bookly.UserStorage) *reservationService {
	return &reservationService{
		reservationStorage: reservations,
		offerStorage:       offers,
		hotelStorage:       hotels,
		reviewStorage:      reviews,
		roomStorage:        rooms,
		userStorage:        users,
	}
}

// CreateReservation handles business logic connected to creating reservations
func (s *reservationService) CreateReservation(ctx context.Context, reservation *bookly.Reservation) error {
	if reservation.ToTime.Before(reservation.FromTime) {
		// Offers is not available during negative time periods
		return bookly.ErrOfferNotAvailable
	}
	if reservation.FromTime.Before(time.Now()) {
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
	rooms, errRooms := s.roomStorage.GetRoomsRelatedWithOffer(ctx, reservation.OfferID)
	if errRooms != nil {
		return errRooms
	}

	for _, room := range rooms {
		owned, errStatus := s.reservationStorage.IsRoomBooked(ctx, room)
		if errStatus != nil {
			return errStatus
		}
		if !owned {
			reservationID, errCreate := s.reservationStorage.CreateReservation(ctx, reservation)
			if errCreate != nil {
				return errCreate
			}
			errLink := s.reservationStorage.CreateReservationRoomLink(ctx, reservationID, room)
			if errLink != nil {
				return errLink
			}
			return nil
		}
	}
	return bookly.ErrNoRoomsLeft
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

		review, err := s.reviewStorage.GetReviewByOwner(ctx, clientID, el.OfferID)
		if err != nil {
			if err == bookly.ErrReviewNotFound {
				obj.Reservation.ReviewID = nil
			} else {
				return nil, err
			}
		} else {
			obj.Reservation.ReviewID = &(review.ID)
		}

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
	if errDelete != nil {
		return errDelete
	}
	errDeleteLinks := s.reservationStorage.DeleteReservationRoomLink(ctx, reservationID)
	return errDeleteLinks
}

// GetHotelReservations retrieves hotel reservations
func (s *reservationService) GetHotelReservations(ctx context.Context, currentOnly bool, hotelID int64, pageNumber int, pageSize int) ([]*bookly.ReservationHotelObject, error) {
	reservations, errGet := s.reservationStorage.GetHotelReservations(ctx, hotelID)
	if errGet != nil {
		return nil, errGet
	}

	list := []*bookly.ReservationHotelObject{}
	for _, el := range reservations {
		if currentOnly && el.ToTime.Before(time.Now()) {
			continue
		}

		obj := &bookly.ReservationHotelObject{}

		client, errGetClient := s.userStorage.GetUser(ctx, el.ClientID)
		if errGetClient != nil {
			return nil, errGetClient
		}

		roomID, errGetRoom := s.reservationStorage.GetRoomFromReservation(ctx, el.ID)
		if errGetRoom != nil {
			return nil, errGetRoom
		}

		roomInfo, errGetRoomInfo := s.roomStorage.GetRoom(ctx, roomID)

		if errGetRoomInfo != nil {
			return nil, errGetRoomInfo
		}

		obj.Room.ID = roomID
		obj.Room.RoomNumber = roomInfo.RoomNumber

		obj.User.ID = el.ClientID
		obj.User.FirstName = client.FirstName
		obj.User.Surname = client.Surname

		obj.Reservation.ReservationID = el.ID
		obj.Reservation.FromTime = el.FromTime
		obj.Reservation.ToTime = el.ToTime
		obj.Reservation.AdultAmount = el.AdultCount
		obj.Reservation.ChildAmount = el.ChildCount

		list = append(list, obj)
	}
	start, end := paging.GetPageItems(pageNumber, pageSize, len(list))
	return list[start:end], nil
}
