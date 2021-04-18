package bookly

import "context"

// Hotel is a domain-level model
type Hotel struct {
	Name        string `json:"hotelName"`
	Description string `json:"hotelDesc"`
	City        string `json:"city"`
	Country     string `json:"country"`
	// todo: change it to different struct as soon as pictures will be implemented
	Pictures       []string `json:"hotelPictures"`
	PreviewPicture string   `json:"hotelPreviewPictures"`
}

// HotelFilter holds information about requested user filtering while fetching hotels
type HotelFilter struct {
	Country   string `json:"country,omitempty"`
	City      string `json:"city,omitempty"`
	HotelName string `json:"hotelName,omitempty"`
}

// HotelListing holds information about hotel previews sent to client
type HotelListing struct {
	HotelID   int64  `json:"hotelId"`
	HotelName string `json:"hotelName"`
	Country   string `json:"country"`
	City      string `json:"city"`
	// todo: change struct when pictures will be implemented
	PreviewPicture string `json:"previewPicture"`
}

// HotelStorage is responsible for operations on hotels
type HotelStorage interface {
	CreateHotel(ctx context.Context, hotel Hotel) (int64, error)
	GetHotelPreviews(ctx context.Context, filter HotelFilter) ([]*HotelListing, error)
	GetHotelDetails(ctx context.Context, hotelID int64) (*Hotel, error)
	UpdateHotelDetails(ctx context.Context, hotelID int64, newHotel Hotel) error
}

// HotelService is responsible for operations on hotels
type HotelService interface {
	GetHotelPreviews(ctx context.Context, filter HotelFilter) ([]*HotelListing, error)
	GetHotelDetails(ctx context.Context, hotelID int64) (*Hotel, error)
	UpdateHotelDetails(ctx context.Context, hotelID int64, newHotel Hotel) error
}
