package main

import (
	"context"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type hotelService struct {
	hotelStorage bookly.HotelStorage
}

func newHotelService(storage bookly.HotelStorage) *hotelService {
	return &hotelService{hotelStorage: storage}
}

// HandleGetHotelDetails handles getting hotel previews
func (os *hotelService) GetHotelPreviews(ctx context.Context, filter bookly.HotelFilter) ([]*bookly.HotelListing, error) {
	listings, err := os.hotelStorage.GetHotelPreviews(ctx, filter)
	if err != nil {
		return nil, err
	}
	return listings, err
}

// HandleGetHotelDetails handles getting hotel details based on its id
func (os *hotelService) GetHotelDetails(ctx context.Context, hotelID int64) (*bookly.Hotel, error) {
	details, err := os.hotelStorage.GetHotelDetails(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	return details, err
}

func validateHotelDetailsRequest(request bookly.Hotel) error {
	if len(request.Name) == 0 {
		return bookly.ErrEmptyHotelName
	}
	return nil
}

// HandleUpdateHotelDetails handles getting hotel details based on its id
func (os *hotelService) UpdateHotelDetails(ctx context.Context, hotelID int64, request bookly.Hotel) error {
	errVal := validateHotelDetailsRequest(request)
	if errVal != nil {
		return errVal
	}
	err := os.hotelStorage.UpdateHotelDetails(ctx, hotelID, request)
	return err
}
