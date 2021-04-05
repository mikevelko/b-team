package main

import (
	"context"
	"fmt"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type offerService struct {
	offerStorage bookly.OfferStorage
}

func newOfferService(storage bookly.OfferStorage) *offerService {
	return &offerService{offerStorage: storage}
}

// IsCreatedOfferValid validates CreateOfferRequest and either returns nil or error with description of wrong parameter
func IsCreatedOfferValid(offer *bookly.Offer) (err error) {
	// cost per adult and per child should not be negative
	if offer.CostPerChild.IsNegative() {
		return fmt.Errorf("offer's cost per child is negative")
	}
	if offer.CostPerAdult.IsNegative() {
		return fmt.Errorf("offer's cost per adult is negative")
	}
	// max guests should be positive
	if offer.MaxGuests <= 0 {
		return fmt.Errorf("offer's max guests number is invalid")
	}
	// offer title should not be empty
	if len(offer.OfferTitle) <= 0 {
		return fmt.Errorf("offer's title is empty")
	}
	// offer should have at least one room connected
	// todo: since we dont have rooms implemented, check that constraint later, when rooms will be implemented
	// todo: also validate pictures once they are implemented
	return nil
}

// HandleCreateOffer handles creating offers from request
func (os *offerService) HandleCreateOffer(ctx context.Context, offer *bookly.Offer) (int64, error) {
	errValidation := IsCreatedOfferValid(offer)
	if errValidation != nil {
		return -1, errValidation
	}

	id, err := os.offerStorage.CreateOffer(ctx, offer, 1)
	if err != nil {
		return -1, err
	}
	return id, err
}

// GetHotelOfferPreviews handles getting offers for particular hotel
func (os *offerService) GetHotelOfferPreviews(ctx context.Context, isActive *bool) ([]bookly.Offer, error) {
	panic("implement me")
}

var _ bookly.OfferService = &offerService{}
