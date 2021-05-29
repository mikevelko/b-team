package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httpapi"
	"github.com/pw-software-engineering/b-team/server/pkg/parse"
	"go.uber.org/zap"
)

type api struct {
	logger             *zap.Logger
	reservationService bookly.ReservationService
}

func newAPI(logger *zap.Logger, service bookly.ReservationService) *api {
	return &api{
		logger:             logger,
		reservationService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/client", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/hotels", func(r chi.Router) {
			r.Post("/{hotelID}/offers/{offerID}/reservations", a.handlePostReservation)
		})
		r.With(auth.SessionMiddleware()).Route("/reservations", func(r chi.Router) {
			r.Delete("/{reservationID}", a.handleDeleteReservation)
			r.Get("/", a.handleGetClientReservations)
		})
	})
}

func (a *api) handlePostReservation(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to add reservation") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest createReservationRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add reservation")
		a.logger.Info("handlePostOffer: could not decode", zap.Error(err))
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	hotelIDStr := chi.URLParam(r, "hotelID")
	hotelID, errConvertHotel := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvertHotel != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvertHotel))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	reservation := decodedRequest.toReservation()
	reservation.ClientID = session.UserID
	reservation.HotelID = hotelID
	reservation.OfferID = offerID
	errAdd := a.reservationService.CreateReservation(r.Context(), reservation)
	//todo: refactor this to use SentinelError and switch to check error type
	if errAdd != nil {
		if errAdd == bookly.ErrOfferNotAvailable {
			httpapi.RespondWithError(w, "Unable to reserve: offer is not available, please refresh offers page.")
			w.WriteHeader(http.StatusBadRequest)
		} else if errAdd == bookly.ErrReservationTooBig {
			httpapi.RespondWithError(w, "Unable to reserve: too much people for this offer.")
			w.WriteHeader(http.StatusBadRequest)
		} else if errAdd == bookly.ErrOfferNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to reserve")
			a.logger.Info("handlePostOffer: could error add reservation", zap.Error(errAdd))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) handleDeleteReservation(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	reservationIDStr := chi.URLParam(r, "reservationID")
	reservationID, errConvertRes := strconv.ParseInt(reservationIDStr, 10, 64)
	if errConvertRes != nil {
		a.logger.Info("Unable to delete reservation, due bad parameter", zap.Error(errConvertRes))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	errDelete := a.reservationService.DeleteReservation(r.Context(), session.UserID, reservationID)
	//todo: refactor this to use SentinelError and switch to check error type
	if errDelete != nil {
		if errDelete == bookly.ErrReservationInProgress {
			httpapi.RespondWithError(w, "Unable to delete reservation: reservation is in progress")
			w.WriteHeader(http.StatusBadRequest)
		} else if errDelete == bookly.ErrReservationNotOwned {
			w.WriteHeader(http.StatusUnauthorized)
		} else if errDelete == bookly.ErrReservationDoesNotExists {
			w.WriteHeader(http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to delete reservation")
			a.logger.Info("handlePostOffer: error could not delete reservation", zap.Error(errDelete))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) handleGetClientReservations(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	pageNumber := parse.IntWithDefault(r.URL.Query().Get("pageNumber"), 1)
	reservesPerPage := parse.IntWithDefault(r.URL.Query().Get("pageSize"), 1)

	reservations, err := a.reservationService.GetClientReservations(r.Context(), session.UserID, pageNumber, reservesPerPage)
	if err != nil {
		a.logger.Info("handlePostOffer: error could not get client reservations", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	httpapi.WriteJSONResponse(a.logger, w, reservations)
}
