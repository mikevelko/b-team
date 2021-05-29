package bookly

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// Offer represents a model for an offer
type Offer struct {
	ID                  int64
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

// OfferClientDetails represents offer details shown to client
type OfferClientDetails struct {
	IsActive     bool            `json:"isActive"`
	IsDeleted    bool            `json:"isDeleted"`
	OfferTitle   string          `json:"offerTitle"`
	CostPerChild decimal.Decimal `json:"costPerChild"`
	CostPerAdult decimal.Decimal `json:"costPerAdult"`
	MaxGuests    int             `json:"maxGuests"`
	Description  string          `json:"offerDescription"`
	Pictures     []*Picture      `json:"offerPictures"`
	// todo: update this as soon as specification will clarify
	AvailabilityTimeIntervals []*interface{} `json:"availabilityTimeIntervals"`
}

// OfferClientFilter holds information about filtering offers for client
type OfferClientFilter struct {
	FromTime  time.Time
	ToTime    time.Time
	MinGuests int
	CostMin   decimal.Decimal
	CostMax   decimal.Decimal
}

// OfferClientPreview holds previews for offers client searched for
type OfferClientPreview struct {
	OfferID      int64           `json:"offerID"`
	OfferTitle   string          `json:"OfferTitle"`
	CostPerChild decimal.Decimal `json:"costPerChild"`
	CostPerAdult decimal.Decimal `json:"costPerAdult"`
	MaxGuests    int             `json:"maxGuests"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	OfferPreviewPicture string `json:"offerPreviewPicture"`
}

// OfferInfoPreview holds info about offer for reservation purpose
type OfferInfoPreview struct {
	OfferID    int64  `json:"offerID"`
	OfferTitle string `json:"offerTitle"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	OfferPreviewPicture string `json:"offerPreviewPicture"`
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
	GetFilteredHotelOfferClientPreviews(ctx context.Context, hotelID int64, filter OfferClientFilter, pageNumber int, itemsPerPage int) ([]*OfferClientPreview, error)
	GetClientHotelOfferDetails(ctx context.Context, hotelID int64, offerID int64) (*OfferClientDetails, error)
}
