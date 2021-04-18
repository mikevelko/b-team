package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"go.uber.org/zap"
)

type api struct {
	logger       *zap.Logger
	hotelService bookly.HotelService
}

func newAPI(logger *zap.Logger, service bookly.HotelService) *api {
	return &api{
		logger:       logger,
		hotelService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/client", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/hotels", func(r chi.Router) {
			r.Get("/", a.handleGetHotelPreviews)
			r.Get("/{hotelID}", a.handleGetHotelDetails)
		})
	})

	router.Route("/api/v1/hotel", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/hotelInfo", func(r chi.Router) {
			r.Get("/", a.handleGetManagerHotelDetails)
			r.Patch("/", a.handlePatchManagerHotelDetails)
		})
	})
}

func (a *api) handleGetHotelPreviews(w http.ResponseWriter, r *http.Request) {
	decodedFilter := bookly.HotelFilter{
		HotelName: r.URL.Query().Get("hotelName"),
		Country:   r.URL.Query().Get("country"),
		City:      r.URL.Query().Get("city"),
	}
	a.logger.Info("Filter values:", zap.String("URL", r.URL.String()), zap.String("HotelName", decodedFilter.HotelName), zap.String("City", decodedFilter.City), zap.String("Country", decodedFilter.Country))
	hotelPreviews, err := a.hotelService.GetHotelPreviews(r.Context(), decodedFilter)
	if err != nil {
		a.logger.Info("Unable to get hotel previews, due to internal error", zap.Error(err))
		httputils.RespondWithError(w, "Unable to get hotel previews")
		return
	}
	httputils.WriteJSONResponse(a.logger, w, hotelPreviews)
	w.WriteHeader(http.StatusOK)
}

func (a *api) handleGetHotelDetails(w http.ResponseWriter, r *http.Request) {
	hotelIDStr := chi.URLParam(r, "hotelID")
	hotelID, errConvert := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get hotel details, due bad parameter", zap.Error(errConvert))
		httputils.RespondWithCode(w, http.StatusNotFound)
		return
	}
	hotel, err := a.hotelService.GetHotelDetails(r.Context(), hotelID)
	if err != nil {
		a.logger.Info("Unable to get hotel details, due to internal error", zap.Error(err))
		httputils.RespondWithCode(w, http.StatusNotFound)
		return
	}
	httputils.WriteJSONResponse(a.logger, w, hotel)
}

func (a *api) handleGetManagerHotelDetails(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	// todo: add role checking for hotel workers
	if session.HotelID == 0 {
		a.logger.Info("Unable to get manager hotel details, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httputils.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	hotel, err := a.hotelService.GetHotelDetails(r.Context(), session.HotelID)
	if err != nil {
		a.logger.Info("Unable to get hotel details, due to internal error", zap.Error(err))
		httputils.RespondWithError(w, "Could not get hotel details")
		return
	}
	httputils.WriteJSONResponse(a.logger, w, hotel)
}

func (a *api) handlePatchManagerHotelDetails(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	// todo: add role checking for hotel workers
	if session.HotelID == 0 {
		a.logger.Info("Unable to get manager hotel details, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httputils.RespondWithError(w, "User is not a manager of any hotel")
		return
	}

	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to update hotel details") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var request HotelPatchRequest
	errDecode := decoder.Decode(&request)
	if errDecode != nil {
		httputils.RespondWithError(w, "Unable to update hotel details")
		a.logger.Info("handlePatchManagerHotelDetails: could not decode", zap.Error(errDecode))
		return
	}

	hotel := bookly.Hotel{
		Name:           request.HotelName,
		Description:    request.HotelDesc,
		PreviewPicture: request.PreviewPicture,
		Pictures:       request.Pictures}
	err := a.hotelService.UpdateHotelDetails(r.Context(), session.HotelID, hotel)
	if err != nil {
		if err == bookly.ErrEmptyHotelName {
			a.logger.Info("Unable to get update hotel details, due to bad request", zap.Error(err))
			httputils.RespondWithError(w, "Could not update hotel details: hotel name is empty")
			return
		}
		httputils.RespondWithError(w, "Unable to update hotel details")
		return
	}
	httputils.RespondWithCode(w, http.StatusOK)
}
