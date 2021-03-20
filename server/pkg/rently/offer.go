package rently

import (
	"context"
	"github.com/shopspring/decimal"
)

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

type OfferStorage interface {
	CreateOffer(ctx context.Context, offer *Offer) (int64, error)
	UpdateOfferStatus(ctx context.Context, offerId int64, isActive bool) error
	GetAllOffers(ctx context.Context, hotelID int) ([]*Offer, error)
}
