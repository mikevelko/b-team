package main

import (
	"context"
	"github.com/pw-software-engineering/b-team/server/pkg/rently"
)

type offerService struct {
	offerStorage rently.OfferStorage
}

func newOfferService(storage rently.OfferStorage) *offerService {
	return &offerService{offerStorage: storage}
}

func (os *offerService) handleCreateOffer(ctx context.Context, request *CreateOfferRequest) (int64, error) {
	var addedOffer rently.Offer
	addedOffer.IsActive = request.Isactive
	addedOffer.CostPerAdult = request.Costperadult
	addedOffer.CostPerChild = request.Costperchild
	addedOffer.Description = request.Description
	addedOffer.MaxGuests = request.Maxguests
	addedOffer.OfferTitle = request.Offertitle
	addedOffer.OfferPreviewPicture = request.Offerpreviewpicture
	//todo: properly handle pictures once they are implemented
	//addedOffer.Pictures = request.Pictures
	addedOffer.Rooms = request.Rooms

	//todo: check CreateRequestOffer data correctness

	id, err := os.offerStorage.CreateOffer(ctx, &addedOffer)
	if err != nil {
		return -1, err
	}
	return id, err
}
