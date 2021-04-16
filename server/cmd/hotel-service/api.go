package main

import (
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
}

func (a *api) handleGetHotelPreviews(w http.ResponseWriter, r *http.Request) {
	decodedFilter := bookly.HotelFilter{
		HotelName: r.URL.Query().Get("hotelName"),
		Country:   r.URL.Query().Get("country"),
		City:      r.URL.Query().Get("city"),
	}
	a.logger.Info("Filter values:", zap.String("URL", r.URL.String()), zap.String("HotelName", decodedFilter.HotelName), zap.String("City", decodedFilter.City), zap.String("Country", decodedFilter.Country))
	hotelPreviews, err := a.hotelService.HandleGetHotelPreviews(r.Context(), decodedFilter)
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
	hotel, err := a.hotelService.HandleGetHotelDetails(r.Context(), hotelID)
	if err != nil {
		a.logger.Info("Unable to get hotel details, due to internal error", zap.Error(err))
		httputils.RespondWithCode(w, http.StatusNotFound)
		return
	}
	httputils.WriteJSONResponse(a.logger, w, hotel)
}
