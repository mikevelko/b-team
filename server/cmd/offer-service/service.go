package main

import (
	"context"
	"fmt"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/paging"
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

// CreateOffer handles creating offers from request
func (os *offerService) CreateOffer(ctx context.Context, hotelID int64, offer *bookly.Offer) (int64, error) {
	errValidation := IsCreatedOfferValid(offer)
	if errValidation != nil {
		return -1, errValidation
	}

	id, err := os.offerStorage.CreateOffer(ctx, offer, hotelID)
	if err != nil {
		return -1, err
	}
	return id, err
}

// GetHotelOfferPreviews handles getting offers for particular hotel
func (os *offerService) GetHotelOfferPreviews(ctx context.Context, hotelID int64, isActive *bool, pageNumber int, itemsPerPage int) ([]*bookly.Offer, error) {
	offers, err := os.offerStorage.GetAllOffers(ctx, hotelID, isActive)
	if err != nil {
		return nil, err
	}
	start, end := paging.GetPageItems(pageNumber, itemsPerPage, len(offers))
	return offers[start:end], err
}

// GetHotelOfferDetails handles getting offer details for particular hotel
func (os *offerService) GetHotelOfferDetails(ctx context.Context, hotelID int64, offerID int64) (*bookly.Offer, error) {
	isOwned, errOwner := os.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if errOwner != nil {
		return nil, errOwner
	}
	if !isOwned {
		return nil, bookly.ErrOfferNotOwned
	}
	offer, err := os.offerStorage.GetSpecificOffer(ctx, offerID)
	return offer, err
}

// MarkHotelOfferAsDeleted implements business logic for marking offer as deleted
func (os *offerService) MarkHotelOfferAsDeleted(ctx context.Context, hotelID int64, offerID int64) error {
	isOwned, errOwned := os.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if errOwned != nil {
		return errOwned
	}
	if !isOwned {
		return bookly.ErrOfferNotOwned
	}
	isActive, errActive := os.offerStorage.IsOfferActive(ctx, offerID)
	if errActive != nil {
		return errActive
	}
	if !isActive {
		return bookly.ErrOfferStillActive
	}
	err := os.offerStorage.SetOfferDeletionStatus(ctx, offerID, true)
	return err
}

// UpdateHotelOffer implements business logic for updating offer details for particular hotel
func (os *offerService) UpdateHotelOffer(ctx context.Context, hotelID int64, offerID int64, offer bookly.Offer) error {
	isOwned, errOwned := os.offerStorage.IsOfferOwnedByHotel(ctx, hotelID, offerID)
	if errOwned != nil {
		return errOwned
	}
	if !isOwned {
		return bookly.ErrOfferNotOwned
	}
	isDeleted, errDeleted := os.offerStorage.IsOfferMarkedAsDeleted(ctx, offerID)
	if errDeleted != nil {
		return errDeleted
	}
	if isDeleted {
		return bookly.ErrOfferDeleted
	}
	errValidation := IsCreatedOfferValid(&offer)
	if errValidation != nil {
		return errValidation
	}
	err := os.offerStorage.UpdateOfferDetails(ctx, offerID, offer)
	return err
}

var _ bookly.OfferService = &offerService{}
