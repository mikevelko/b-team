package bookly

import (
	"context"

	"github.com/shopspring/decimal"
)

// Offer represents a model for an offer
type Offer struct {
	IsActive            bool
	OfferTitle          string
	CostPerChild        decimal.Decimal
	CostPerAdult        decimal.Decimal
	MaxGuests           int
	Description         string
	OfferPreviewPicture string
	Pictures            []*Picture
	Rooms               []*Room
}

// OfferStorage is responsible for persisting offers
type OfferStorage interface {
	CreateOffer(ctx context.Context, offer *Offer, hotelID int64) (int64, error)
	SetOfferDeletionStatus(ctx context.Context, offerID int64, isDeleted bool) error
	GetAllOffers(ctx context.Context, hotelID int64, isActive *bool) ([]*Offer, error)
	GetSpecificOffer(ctx context.Context, offerID int64) (*Offer, error)
	UpdateOfferDetails(ctx context.Context, offerID int64, newOffer Offer) error
	IsOfferActive(ctx context.Context, offerID int64) (bool, error)
	IsOfferOwnedByHotel(ctx context.Context, hotelID int64, offerID int64) (bool, error)
	IsOfferMarkedAsDeleted(ctx context.Context, offerID int64) (bool, error)
}

// OfferService is a service which is responsible for actions related to offers
type OfferService interface {
	CreateOffer(ctx context.Context, hotelID int64, offer *Offer) (int64, error)
	GetHotelOfferPreviews(ctx context.Context, hotelID int64, isActive *bool, pageNumber int, itemsPerPage int) ([]*Offer, error)
	GetHotelOfferDetails(ctx context.Context, hotelID int64, offerID int64) (*Offer, error)
	UpdateHotelOffer(ctx context.Context, hotelID int64, offerID int64, offer Offer) error
	MarkHotelOfferAsDeleted(ctx context.Context, hotelID int64, offerID int64) error
}
