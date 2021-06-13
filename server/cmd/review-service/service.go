package main

import (
	"context"
	"time"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type reviewService struct {
	reviewStorage      bookly.ReviewStorage
	offerStorage       bookly.OfferStorage
	userStorage        bookly.UserStorage
	reservationStorage bookly.ReservationStorage
}

func newReviewService(reviewStorage bookly.ReviewStorage, offerStorage bookly.OfferStorage, userStorage bookly.UserStorage, reservationStorage bookly.ReservationStorage) *reviewService {
	return &reviewService{
		reviewStorage:      reviewStorage,
		offerStorage:       offerStorage,
		userStorage:        userStorage,
		reservationStorage: reservationStorage,
	}
}

var _ bookly.ReviewService = &reviewService{}

// GetReviewsOfHotel implements business logic of reviews
func (r *reviewService) GetReviewsOfHotel(ctx context.Context, hotelID int64, userID int64) ([]*bookly.Review, error) {
	offers, err := r.offerStorage.GetAllOffers(ctx, hotelID, nil)
	if nil != err {
		return nil, err
	}
	reviews := []*bookly.Review{}
	for _, offer := range offers {
		offerReviews, err := r.GetReviewsOfOffer(ctx, offer.ID, userID)
		if nil != err {
			return nil, err
		}
		for _, review := range offerReviews {
			reviews = append(reviews, review)
		}
	}
	return reviews, nil
}

// GetReviewsOfOffer implements business logic of reviews
func (r *reviewService) GetReviewsOfOffer(ctx context.Context, offerID int64, userID int64) ([]*bookly.Review, error) {
	reviews, err := r.reviewStorage.GetReviewsOfOffer(ctx, offerID)
	if err != nil {
		return nil, err
	}
	for _, review := range reviews {
		user, err := r.userStorage.GetUser(ctx, review.UserID)
		if err != nil {
			return nil, err
		}
		review.ReviewerUsername = user.UserName
	}
	return reviews, nil
}

// GetReviewsOfReservation implements business logic of reviews
func (r *reviewService) GetReviewsOfReservation(ctx context.Context, reservationID int64, userID int64) (bookly.Review, error) {
	offerID, err := r.getOfferID(ctx, userID, reservationID)
	if err != nil {
		return bookly.Review{}, err
	}

	review, err := r.reviewStorage.GetReviewByOwner(ctx, userID, offerID)
	if err != nil {
		return bookly.Review{}, err
	}

	user, err := r.userStorage.GetUser(ctx, review.UserID)
	if err != nil {
		return bookly.Review{}, err
	}
	review.ReviewerUsername = user.UserName

	return review, nil
}

// CreateOrUpdateReview implements business logic of reviews
func (r *reviewService) CreateOrUpdateReview(ctx context.Context, review bookly.Review, userID int64, reservationID int64) (int64, error) {
	offerID, err := r.getOfferID(ctx, userID, reservationID)
	if err != nil {
		return 0, err
	}

	err = r.reviewStorage.DeleteReviewByOwner(ctx, userID, offerID)
	if err != nil && err != bookly.ErrReviewNotFound {
		return 0, err
	}

	review.OfferID = offerID
	review.UserID = userID
	review.ReviewDate = time.Now()

	reviewID, err := r.reviewStorage.CreateReview(ctx, review)
	if err != nil && err != bookly.ErrReviewNotFound {
		return 0, err
	}
	return reviewID, nil
}

// DeleteReview implements business logic of reviews
func (r *reviewService) DeleteReview(ctx context.Context, userID int64, reservationID int64) error {
	offerID, err := r.getOfferID(ctx, userID, reservationID)
	if err != nil {
		return err
	}

	err = r.reviewStorage.DeleteReviewByOwner(ctx, userID, offerID)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewService) getOfferID(ctx context.Context, userID int64, reservationID int64) (int64, error) {
	isOwned, err := r.reservationStorage.IsReservationOwnedByClient(ctx, userID, reservationID)
	if err != nil {
		return 0, err
	}
	if isOwned == false {
		return 0, bookly.ErrReviewNotOwned
	}
	reservation, err := r.reservationStorage.GetSpecificReservation(ctx, reservationID)
	if err != nil {
		return 0, err
	}

	return reservation.OfferID, nil
}
