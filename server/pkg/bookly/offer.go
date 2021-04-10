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
	CreateOffer(ctx context.Context, offer *Offer, hotelID int) (int64, error)
	UpdateOfferStatus(ctx context.Context, offerID int64, isActive bool) error
	GetAllOffers(ctx context.Context, hotelID int, isActive *bool) ([]*Offer, error)
}

// OfferService is a service which is responsible for actions related to offers
type OfferService interface {
	HandleCreateOffer(ctx context.Context, offer *Offer) (int64, error)
	GetHotelOfferPreviews(ctx context.Context, isActive *bool) ([]*Offer, error)
}
