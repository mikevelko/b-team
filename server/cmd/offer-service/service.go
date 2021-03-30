package main

import (
	"context"
	"fmt"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type offerService struct {
	offerStorage bookly.OfferStorage
}

func newOfferService(storage bookly.OfferStorage) *offerService {
	return &offerService{offerStorage: storage}
}

//IsCreatedOfferValid validates CreateOfferRequest and either returns nil or error with description of wrong parameter
func IsCreatedOfferValid(offer *CreateOfferRequest) (err error) {
	//cost per adult and per child should not be negative
	if offer.Costperchild.IsNegative() {
		return fmt.Errorf("offer's cost per child is negative")
	}
	if offer.Costperadult.IsNegative() {
		return fmt.Errorf("offer's cost per adult is negative")
	}
	//max guests should be positive
	if offer.Maxguests <= 0 {
		return fmt.Errorf("offer's max guests number is invalid")
	}
	//offer title should not be empty
	if len(offer.Offertitle) <= 0 {
		return fmt.Errorf("offer's title is empty")
	}
	//offer should have at least one room connected
	//todo: since we dont have rooms implemented, check that constraint later, when rooms will be implemented
	//todo: also validate pictures once they are implemented
	return nil
}

//handleCreateOffer validates Create Offer Request and passes arguments to business logic
func (os *offerService) handleCreateOffer(ctx context.Context, request *CreateOfferRequest, hotelToken string) (int64, error) {

	errValidation := IsCreatedOfferValid(request)
	if errValidation != nil {
		return -1, errValidation
	}

	var addedOffer bookly.Offer
	addedOffer.IsActive = request.Isactive
	addedOffer.CostPerAdult = request.Costperadult
	addedOffer.CostPerChild = request.Costperchild
	addedOffer.Description = request.Description
	addedOffer.MaxGuests = request.Maxguests
	addedOffer.OfferTitle = request.Offertitle
	addedOffer.OfferPreviewPicture = request.Offerpreviewpicture
	// todo: properly handle pictures once they are implemented
	// addedOffer.Pictures = request.Pictures
	addedOffer.Rooms = request.Rooms

	id, err := os.offerStorage.CreateOffer(ctx, &addedOffer, 1)
	if err != nil {
		return -1, err
	}
	return id, err
}
