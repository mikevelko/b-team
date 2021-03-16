package rently

type Offer struct {
	Id int
}

type OfferStorage interface {
	DeleteOffer(id int) error
	GetAllOffers(hotelID int) ([]*Offer, error)
}
