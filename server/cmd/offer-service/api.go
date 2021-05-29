package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/pw-software-engineering/b-team/server/pkg/parse"
	"github.com/shopspring/decimal"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httpapi"
	"go.uber.org/zap"
)

type api struct {
	logger       *zap.Logger
	offerService bookly.OfferService
}

func newAPI(logger *zap.Logger, service bookly.OfferService) *api {
	return &api{
		logger:       logger,
		offerService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/hotel", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/offers", func(r chi.Router) {
			r.Post("/", a.handlePostOffer)
			r.Get("/", a.handleGetOffers)
			r.Get("/{offerID}", a.handleGetOfferDetails)
			r.Delete("/{offerID}", a.handleDeleteOffer)
			r.Patch("/{offerID}", a.handleUpdateOfferDetails)
		})
	})
	router.Route("/api/v1/client", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/hotels", func(r chi.Router) {
			r.Get("/{hotelID}/offers", a.handleGetClientHotelOffers)
			r.Get("/{hotelID}/offers/{offerID}", a.handleGetClientHotelOfferDetails)
		})
	})
}

func (a *api) handleGetClientHotelOffers(w http.ResponseWriter, r *http.Request) {
	filter := bookly.OfferClientFilter{}
	filter.CostMin = parse.DecimalWithDefault(r.URL.Query().Get("costMin"), decimal.NewFromInt(0))
	filter.CostMax = parse.DecimalWithDefault(r.URL.Query().Get("costMax"), decimal.NewFromInt(2100000000))
	filter.MinGuests = parse.IntWithDefault(r.URL.Query().Get("minGuests"), 1)

	hotelIDStr := chi.URLParam(r, "hotelID")
	hotelID, errConvert := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get client hotel offers, due to bad hotelID parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusBadRequest)
		return
	}

	pageNumber := parse.IntWithDefault(r.URL.Query().Get("pageNumber"), 1)
	offersPerPage := parse.IntWithDefault(r.URL.Query().Get("pageSize"), 1)
	a.logger.Info("Get Offers", zap.Int("pageNumber", pageNumber), zap.Int("itemsPerPage", offersPerPage))
	offerPreviews, err := a.offerService.GetFilteredHotelOfferClientPreviews(r.Context(), hotelID, filter, pageNumber, offersPerPage)
	if err != nil {
		a.logger.Info("Unable to get client hotel offers, due to internal error", zap.Error(err))
		httpapi.RespondWithError(w, "Unable to get offers")
		return
	}
	httpapi.WriteJSONResponse(a.logger, w, offerPreviews)
}

func (a *api) handleGetClientHotelOfferDetails(w http.ResponseWriter, r *http.Request) {
	hotelIDStr := chi.URLParam(r, "hotelID")
	hotelID, errConvert := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get client hotel offer details, due to bad hotelID parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusBadRequest)
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvertOf := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvertOf != nil {
		a.logger.Info("Unable to get client hotel offer details, due to bad offerID parameter", zap.Error(errConvertOf))
		httpapi.RespondWithCode(w, http.StatusBadRequest)
		return
	}
	result, err := a.offerService.GetClientHotelOfferDetails(r.Context(), hotelID, offerID)
	if err != nil {
		a.logger.Info("Unable to get client hotel offer details, due to internal error", zap.Error(err))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	httpapi.WriteJSONResponse(a.logger, w, result)
}

func (a *api) handlePostOffer(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to post offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to add offer") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest createOfferRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add offer")
		a.logger.Info("handlePostOffer: could not decode", zap.Error(err))
		return
	}

	id, err := a.offerService.CreateOffer(r.Context(), session.HotelID, decodedRequest.toOffer())
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add offer")
		a.logger.Info("handlePostOffer: could error create offer", zap.Error(err))
		return
	}
	idResponse := &offerIDResponse{id}
	httpapi.WriteJSONResponse(a.logger, w, idResponse)
}

func (a *api) handleGetOffers(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to get offers, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	pageNumberStr := r.URL.Query().Get("pageNumber")
	offersPerPageStr := r.URL.Query().Get("pageSize")
	isActiveStr := r.URL.Query().Get("isActive")
	pageNumber, errPN := strconv.ParseInt(pageNumberStr, 10, 32)
	if errPN != nil {
		pageNumber = 1
	}
	offersPerPage, errOPP := strconv.ParseInt(offersPerPageStr, 10, 32)
	if errOPP != nil {
		offersPerPage = 1
	}
	var isActive *bool
	isActiveVal, errBool := strconv.ParseBool(isActiveStr)
	if errBool != nil {
		isActive = nil
	} else {
		isActive = &isActiveVal
	}
	a.logger.Info("Get Offers", zap.Int64("pageNumber", pageNumber), zap.Int64("itemsPerPage", offersPerPage))
	// todo: supply proper hotel token once hotel verification will be available
	offerPreviews, err := a.offerService.GetHotelOfferPreviews(r.Context(), session.HotelID, isActive, int(pageNumber), int(offersPerPage))
	if err != nil {
		a.logger.Info("Unable to get offers, due to internal error", zap.Error(err))
		httpapi.RespondWithError(w, "Unable to get offers")
		return
	}
	offersResponse := &getOffersResponse{Offers: offerPreviewsFromOffers(offerPreviews)}
	httpapi.WriteJSONResponse(a.logger, w, offersResponse)
}

func (a *api) handleGetOfferDetails(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to get offer details, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get offers details, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	offer, err := a.offerService.GetHotelOfferDetails(r.Context(), session.HotelID, offerID)
	if err != nil {
		a.logger.Info("Unable to get offers details, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		}
		return
	}

	response := &offerDetailsResponse{
		IsActive:            offer.IsActive,
		OfferTitle:          offer.OfferTitle,
		CostPerAdult:        parse.DecimalToFloat(offer.CostPerAdult),
		CostPerChild:        parse.DecimalToFloat(offer.CostPerChild),
		MaxGuests:           offer.MaxGuests,
		Description:         offer.Description,
		OfferPreviewPicture: offer.OfferPreviewPicture,
	}
	httpapi.WriteJSONResponse(a.logger, w, response)
}

func (a *api) handleDeleteOffer(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to delete offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	err := a.offerService.MarkHotelOfferAsDeleted(r.Context(), session.HotelID, offerID)
	if err != nil {
		a.logger.Info("Unable to delete offer, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else if errors.Is(err, bookly.ErrOfferStillActive) {
			httpapi.RespondWithError(w, "Offer is still active")
			httpapi.RespondWithCode(w, http.StatusConflict)
		} else {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		}
		return
	}
	httpapi.RespondWithCode(w, http.StatusOK)
}

func (a *api) handleUpdateOfferDetails(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to update offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to update offer") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest createOfferRequest
	errDecode := decoder.Decode(&decodedRequest)
	if errDecode != nil {
		httpapi.RespondWithError(w, "Unable to update offer")
		a.logger.Info("handleUpdateOffer: could not decode", zap.Error(errDecode))
		return
	}
	err := a.offerService.UpdateHotelOffer(r.Context(), session.HotelID, offerID, *decodedRequest.toOffer())
	if err != nil {
		a.logger.Info("Unable to update offer, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		}
		return
	}
	httpapi.RespondWithCode(w, http.StatusOK)
}
