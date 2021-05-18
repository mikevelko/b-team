package main

import (
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/parse"
	"github.com/shopspring/decimal"
)

type offerIDResponse struct {
	OfferID int64 `json:"offerID"`
}

// createOfferRequest represents deserialized data from CreateOffer request
type createOfferRequest struct {
	IsActive     bool            `json:"isActive"`
	OfferTitle   string          `json:"offerTitle"`
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
	IsActive     bool    `json:"isActive"`
	OfferTitle   string  `json:"offerTitle"`
	CostPerChild float64 `json:"costPerChild"`
	CostPerAdult float64 `json:"costPerAdult"`
	MaxGuests    int     `json:"maxGuests"`
	Description  string  `json:"description"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	OfferPreviewPicture string        `json:"offerPreviewPicture"`
	Pictures            []interface{} `json:"pictures"`
}

// getOffersResponse represents serialized offers sent back to request
type getOffersResponse struct {
	Offers []*offerPreview `json:"offerPreview"`
}

func offerPreviewFromOffer(o *bookly.Offer) *offerPreview {
	return &offerPreview{
		OfferID:             o.ID,
		IsActive:            o.IsActive,
		OfferTitle:          o.OfferTitle,
		CostPerChild:        parse.DecimalToFloat(o.CostPerChild),
		CostPerAdult:        parse.DecimalToFloat(o.CostPerAdult),
		MaxGuests:           o.MaxGuests,
		OfferPreviewPicture: o.OfferPreviewPicture,
	}
}

func offerPreviewsFromOffers(offers []*bookly.Offer) []*offerPreview {
	ops := make([]*offerPreview, len(offers))
	for i, offer := range offers {
		ops[i] = offerPreviewFromOffer(offer)
	}
	return ops
}

type offerPreview struct {
	OfferID             int64   `json:"offerID"`
	IsActive            bool    `json:"isActive"`
	OfferTitle          string  `json:"offerTitle"`
	CostPerChild        float64 `json:"costPerChild"`
	CostPerAdult        float64 `json:"costPerAdult"`
	MaxGuests           int     `json:"maxGuests"`
	OfferPreviewPicture string  `json:"offerPreviewPicture"`
}

// toOffer maps a request to add an offer to model offer
func (c createOfferRequest) toOffer() *bookly.Offer {
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
