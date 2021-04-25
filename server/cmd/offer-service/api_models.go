package main

import (
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/shopspring/decimal"
)

type offerIDResponse struct {
	OfferID int64 `json:"offerID"`
}

// CreateOfferRequest represents deserialized data from CreateOffer request
type CreateOfferRequest struct {
	IsActive     bool            `json:"IsActive"`
	OfferTitle   string          `json:"OfferTitle"`
	CostPerChild decimal.Decimal `json:"costPerChild"`
	CostPerAdult decimal.Decimal `json:"costPerAdult"`
	MaxGuests    int             `json:"maxGuests"`
	Description  string          `json:"description"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	OfferPreviewPicture string        `json:"offerPreviewPicture"`
	Pictures            []interface{} `json:"pictures"`
	Rooms               []string      `json:"rooms"`
}

type offerDetailsResponse struct {
	IsActive     bool            `json:"IsActive"`
	OfferTitle   string          `json:"OfferTitle"`
	CostPerChild decimal.Decimal `json:"costPerChild"`
	CostPerAdult decimal.Decimal `json:"costPerAdult"`
	MaxGuests    int             `json:"maxGuests"`
	Description  string          `json:"description"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	OfferPreviewPicture string        `json:"offerPreviewPicture"`
	Pictures            []interface{} `json:"pictures"`
}

// GetOffersResponse represents serialized offers sent back to request
type GetOffersResponse struct {
	Offers []*bookly.Offer `json:"offerPreview"`
	// In case of questions: marshal replace pointers to structures
	// Example: https://play.golang.org/p/NkPCH21ukj
}

// ToOffer maps a request to add an offer to model offer
func (c CreateOfferRequest) ToOffer() *bookly.Offer {
	return &bookly.Offer{
		IsActive:            c.IsActive,
		OfferTitle:          c.OfferTitle,
		CostPerChild:        c.CostPerAdult,
		CostPerAdult:        c.CostPerChild,
		MaxGuests:           c.MaxGuests,
		Description:         c.Description,
		OfferPreviewPicture: c.OfferPreviewPicture,
	}
}
