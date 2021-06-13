package main

import (
	//"encoding/json"
	//"net/http"

	//"github.com/pw-software-engineering/b-team/server/pkg/auth"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httpapi"

	//"github.com/pw-software-engineering/b-team/server/pkg/httpapi"
	"go.uber.org/zap"
)

type api struct {
	logger        *zap.Logger
	reviewService bookly.ReviewService
}

func newAPI(logger *zap.Logger, service bookly.ReviewService) *api {
	return &api{
		logger:        logger,
		reviewService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/client", func(r chi.Router) {
			r.Get("/hotels/{hotelID}/reviews", a.handleGetHotelReviews)
			r.Get("/hotels/{hotelID}/offers/{offerID}/reviews", a.handleGetOfferReviews)
			r.Get("/reservations/{reservationID}/review", a.handleGetReview)
			r.Delete("/reservations/{reservationID}/review", a.handleDeleteReview)
			r.Put("/reservations/{reservationID}/review", a.handleUpdateReview)
		})
	})
}

func (a *api) handleGetHotelReviews(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())

	hotelIDStr := chi.URLParam(r, "hotelID")
	hotelID, errConvertHotel := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvertHotel != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvertHotel))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	reviews, err := a.reviewService.GetReviewsOfHotel(r.Context(), hotelID, session.UserID)
	if err != nil {
		httpapi.RespondWithErrorCode(w, "Unable to reserve: reviews is not available.", http.StatusNotFound)
		return
	}

	httpapi.WriteJSONResponse(a.logger, w, reviews)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handleGetOfferReviews(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())

	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	hotelIDStr := chi.URLParam(r, "hotelID")
	_, errConvertHotel := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvertHotel != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvertHotel))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	reviews, err := a.reviewService.GetReviewsOfOffer(r.Context(), offerID, session.UserID)
	if err != nil {
		httpapi.RespondWithErrorCode(w, "Unable to reserve: reviews is not available.", http.StatusNotFound)
		return
	}

	httpapi.WriteJSONResponse(a.logger, w, reviews)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handleGetReview(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())

	reservationIDStr := chi.URLParam(r, "reservationID")
	reservationID, errConvert := strconv.ParseInt(reservationIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get review, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	review, err := a.reviewService.GetReviewsOfReservation(r.Context(), reservationID, session.UserID)
	if err != nil {
		if err == bookly.ErrReviewNotFound {
			httpapi.RespondWithErrorCode(w, "Unable to get review: no review.", http.StatusNotFound)
			return
		}
		if err == bookly.ErrReviewNotOwned {
			httpapi.RespondWithErrorCode(w, "Unable to get review: no owned.", http.StatusNotFound)
			return
		}
		httpapi.RespondWithError(w, "Unable to reserve")
		a.logger.Info("handlePostOffer: could error add reservation", zap.Error(err))
		return
	}
	httpapi.WriteJSONResponse(a.logger, w, review)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handleDeleteReview(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())

	reservationIDStr := chi.URLParam(r, "reservationID")
	reservationID, errConvert := strconv.ParseInt(reservationIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get review, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	err := a.reviewService.DeleteReview(r.Context(), session.UserID, reservationID)
	if err != nil {
		if err == bookly.ErrReviewNotFound {
			httpapi.RespondWithErrorCode(w, "Unable to delete review: no review.", http.StatusNotFound)
			return
		}
		if err == bookly.ErrReviewNotOwned {
			httpapi.RespondWithErrorCode(w, "Unable to delete review: no owned.", http.StatusNotFound)
			return
		}
		httpapi.RespondWithError(w, "Unable to delete: reviews is not available.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handleUpdateReview(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to add review") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest createReviewModel
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add review")
		a.logger.Info("handlePostReview: could not decode", zap.Error(err))
		return
	}

	reservationIDStr := chi.URLParam(r, "reservationID")
	reservationID, errConvert := strconv.ParseInt(reservationIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get review, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	review := bookly.Review{
		Content: decodedRequest.Content,
		Rating:  decodedRequest.Rating,
	}

	_, err = a.reviewService.CreateOrUpdateReview(r.Context(), review, session.UserID, reservationID)

	if err != nil {
		if err == bookly.ErrReviewNotFound {
			httpapi.RespondWithErrorCode(w, "Unable to put review: no review.", http.StatusNotFound)
			return
		}
		if err == bookly.ErrReviewNotOwned {
			httpapi.RespondWithErrorCode(w, "Unable to put review: no owned.", http.StatusNotFound)
			return
		}
		httpapi.RespondWithError(w, "Unable to reserve")
		a.logger.Info("handlePostReview: could error add review", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
