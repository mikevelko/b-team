package main

import (
	"context"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type reviewService struct {
	reviewStorage bookly.ReviewStorage
}

func newReviewService(reviewStorage bookly.ReviewStorage) *reviewService {
	return &reviewService{
		reviewStorage: reviewStorage,
	}
}

var _ bookly.ReviewService = &reviewService{}

func (r reviewService) GetReviewsOfOffer(ctx context.Context, offerID int64, userID int64) ([]*bookly.Review, error) {
	panic("implement me")
}

func (r reviewService) CreateReview(ctx context.Context, review bookly.Review, userID int64, offerID int64) (int64, error) {
	panic("implement me")
}

func (r reviewService) DeleteReview(ctx context.Context, reviewID int64, userID int64, offerID int64) error {
	panic("implement me")
}

func (r reviewService) UpdateReview(ctx context.Context, review bookly.Review, userID int64, offerID int64) (int64, error) {
	panic("implement me")
}
