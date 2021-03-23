package rently

import (
	"context"

	"github.com/shopspring/decimal"
)

//Offer represents a model for an offer
type Offer struct {
	IsActive            bool            `json:"isActive"`
	OfferTitle          string          `json:"offerTitle"`
	CostPerChild        decimal.Decimal `json:"costPerChild"`
	CostPerAdult        decimal.Decimal `json:"costPerAdult"`
	MaxGuests           int             `json:"maxGuests"`
	Description         string          `json:"description"`
	OfferPreviewPicture string          `json:"offerPreviewPicture"`
	Pictures            []*Picture      `json:"pictures"`
	Rooms               []string        `json:"rooms"`
}

//OfferStorage is responsible for persisting offers
type OfferStorage interface {
	CreateOffer(ctx context.Context, offer *Offer) (int64, error)
	UpdateOfferStatus(ctx context.Context, offerID int64, isActive bool) error
	GetAllOffers(ctx context.Context, hotelID int) ([]*Offer, error)
}
