package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type api struct {
	logger       *zap.Logger
	offerService bookly.OfferService
}

//todo: those structures below should be moved to other go file for readibility

// CreateOfferRequest represents deserialized data from CreateOffer request
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

//GetOffersRequest represents deserialized data from GetOffers Request
type GetOffersRequest struct {
	Isactive *bool `json:"isActive,omitempty"`
}

//GetOffersResponse represents serialized offers sent back to request
type GetOffersResponse struct {
	Offers []bookly.Offer `json:"offerPreview"`
}

// ToOffer maps a request to add an offer to model offer
func (c CreateOfferRequest) ToOffer() *bookly.Offer {
	return &bookly.Offer{
		IsActive:            c.Isactive,
		OfferTitle:          c.Offertitle,
		CostPerChild:        c.Costperadult,
		CostPerAdult:        c.Costperchild,
		MaxGuests:           c.Maxguests,
		Description:         c.Description,
		OfferPreviewPicture: c.Offerpreviewpicture,
	}
}

type offerIDResponse struct {
	offerID int64
}

func newAPI(logger *zap.Logger, service bookly.OfferService) *api {
	return &api{
		logger:       logger,
		offerService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/hotel", func(r chi.Router) {
		r.Route("/offers", func(r chi.Router) {
			r.Post("/", a.handlePostOffer)
			r.Get("/", a.handleGetOffers)
		})
	})
}

func (a *api) handlePostOffer(w http.ResponseWriter, r *http.Request) {
	// hotelToken := r.Header.Get("x-hotel-token")
	// todo: check x-hotel-token for correctness

	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to add offer") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest CreateOfferRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httputils.RespondWithError(w, "Unable to add offer")
		return
	}
	// todo: supply proper hotel token once hotel verification will be available
	id, err := a.offerService.HandleCreateOffer(r.Context(), decodedRequest.ToOffer(), "")
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
	httputils.WriteJSONResponse(w, js)
}

func (a *api) handleGetOffers(w http.ResponseWriter, r *http.Request) {
	// hotelToken := r.Header.Get("x-hotel-token")
	// todo: check x-hotel-token for correctness

	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to get offers") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest GetOffersRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httputils.RespondWithError(w, "Unable to get offers")
		return
	}
	// todo: supply proper hotel token once hotel verification will be available
	offerPreviews, err := a.offerService.GetHotelOfferPreviews(r.Context(), decodedRequest.Isactive, "")
	if err != nil {
		httputils.RespondWithError(w, "Unable to get offers")
		return
	}
	offersResponse := GetOffersResponse{Offers: offerPreviews}
	js, err := json.Marshal(offersResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSONResponse(w, js)
}
