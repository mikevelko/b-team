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
	router.Route("/api/v1/client", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/hotels", func(r chi.Router) {
			r.Get("/{hotelID}/offers/{offerID}/reviews", a.handleGetOfferReviews)
			r.Post("/{hotelID}/offers/{offerID}/reviews", a.handlePostReview)
			r.Delete("/{hotelID}/offers/{offerID}/reviews/{reviewID}", a.handleDeleteReview)
			r.Put("/{hotelID}/offers/{offerID}/reviews/{reviewID}", a.handleUpdateReview)
		})
	})
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
		httpapi.RespondWithError(w, "Unable to reserve: reviews is not available.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	httpapi.WriteJSONResponse(a.logger, w, reviews)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handlePostReview(w http.ResponseWriter, r *http.Request) {
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

	review := bookly.Review{
		Content: decodedRequest.Content,
		Rating:  decodedRequest.Rating,
	}

	reviewID, err := a.reviewService.CreateReview(r.Context(), review, session.UserID, offerID)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to reserve")
		a.logger.Info("handlePostReview: could error add review", zap.Error(err))
		return
	}

	respond := createReviewRespondModel{
		ReviewID: reviewID,
	}

	httpapi.WriteJSONResponse(a.logger, w, respond)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *api) handleDeleteReview(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())

	reviewIDStr := chi.URLParam(r, "reviewID")
	reviewID, errConvertReview := strconv.ParseInt(reviewIDStr, 10, 64)
	if errConvertReview != nil {
		a.logger.Info("Unable to delete review, due bad parameter", zap.Error(errConvertReview))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete review, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}
	hotelIDStr := chi.URLParam(r, "hotelID")
	_, errConvertHotel := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvertHotel != nil {
		a.logger.Info("Unable to delete review, due bad parameter", zap.Error(errConvertHotel))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	err := a.reviewService.DeleteReview(r.Context(), reviewID, session.UserID, offerID)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to reserve: reviews is not available.")
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
	var decodedRequest bookly.Review
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add review")
		a.logger.Info("handlePostReview: could not decode", zap.Error(err))
		return
	}

	reviewIDStr := chi.URLParam(r, "reviewID")
	reviewID, errConvertReview := strconv.ParseInt(reviewIDStr, 10, 64)
	if errConvertReview != nil {
		a.logger.Info("Unable to put review, due bad parameter", zap.Error(errConvertReview))
		httpapi.RespondWithCode(w, http.StatusNotFound)
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
	_, errConvertHotel := strconv.ParseInt(hotelIDStr, 10, 64)
	if errConvertHotel != nil {
		a.logger.Info("Unable to update offer, due bad parameter", zap.Error(errConvertHotel))
		httpapi.RespondWithCode(w, http.StatusNotFound)
		return
	}

	decodedRequest.ID = reviewID
	_, err = a.reviewService.UpdateReview(r.Context(), decodedRequest, session.UserID, offerID)

	if err != nil {
		httpapi.RespondWithError(w, "Unable to reserve")
		a.logger.Info("handlePostReview: could error add review", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
