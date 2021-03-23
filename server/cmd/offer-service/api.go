package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type api struct {
	logger       *zap.Logger
	offerService *offerService
}

//CreateOfferRequest represents deserialized data from CreateOffer request
type CreateOfferRequest struct {
	Isactive     bool            `json:"isActive"`
	Offertitle   string          `json:"offerTitle"`
	Costperchild decimal.Decimal `json:"costPerChild"`
	Costperadult decimal.Decimal `json:"costPerAdult"`
	Maxguests    int             `json:"maxGuests"`
	Description  string          `json:"description"`
	// todo: pictures will probably need different struct, reimplement it as soon as they are implemented
	Offerpreviewpicture string        `json:"offerPreviewPicture"`
	Pictures            []interface{} `json:"pictures"`
	Rooms               []string      `json:"rooms"`
}

type offerIDResponse struct {
	offerID int64
}

func newAPI(logger *zap.Logger, service *offerService) *api {
	return &api{
		logger:       logger,
		offerService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/hotel", func(r chi.Router) {
		r.Route("/offers", func(r chi.Router) {
			r.Post("/", a.handlePostOffer)
		})
	})
}

func (a *api) handlePostOffer(w http.ResponseWriter, r *http.Request) {
	// hotelToken := r.Header.Get("x-hotel-token")
	// todo: check x-hotel-token for correctness

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		httputils.RespondWithError(w, "Unable to add offer")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decoded CreateOfferRequest
	err := decoder.Decode(&decoded)
	if err != nil {
		httputils.RespondWithError(w, "Unable to add offer")
		return
	}
	id, err := a.offerService.handleCreateOffer(r.Context(), &decoded)
	if err != nil {
		httputils.RespondWithError(w, "Unable to add offer")
		return
	}
	idResponse := offerIDResponse{id}
	js, err := json.Marshal(idResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// todo: tests for this endpoint if applicable
}
