package main

// HotelPatchRequest holds info about new hotel details requested by manager
type HotelPatchRequest struct {
	HotelName string `json:"hotelName,omitempty"`
	HotelDesc string `json:"hotelDesc,omitempty"`
	// todo: change it to different struct as soon as pictures will be implemented
	PreviewPicture string   `json:"hotelPreviewPicture,omitempty"`
	Pictures       []string `json:"hotelPictures,omitempty"`
}
