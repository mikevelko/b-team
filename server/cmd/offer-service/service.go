package main

import (
	"github.com/pw-software-engineering/b-team/server/pkg/rently"
	"net/http"
)

type OfferService struct {
	offerStorage rently.OfferStorage
}

func (o OfferService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
