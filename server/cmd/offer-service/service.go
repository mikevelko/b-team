package main

import (
	"context"
	"github.com/pw-software-engineering/b-team/server/pkg/rently"
)

type offerService struct {
	offerStorage rently.OfferStorage
}

func newOfferService(storage rently.OfferStorage) *offerService{
	return &offerService{offerStorage: storage}
}

func (os *offerService) handleCreateOffer(ctx context.Context, request *CreateOfferRequest) (int, error){
	//todo: add offer creation handling here
	panic("implement me")
}

