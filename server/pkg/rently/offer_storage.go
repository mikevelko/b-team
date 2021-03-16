package rently

type OfferStorage interface{
    GetOffers() error
}
