package bookly

import (
	"context"
	"time"
)

// Review is a model for room, not implemented yed
type Review struct {
	ID         int64
	OfferID    int64
	UserID     int64
	Content    string
	Rating     int32 //[1-5]
	ReviewDate time.Time
}

// ReviewStorage is responsible for persisting review
type ReviewStorage interface {
	CreateReview(ctx context.Context, review Review) (int64, error)
	DeleteReview(ctx context.Context, reviewID int64) error
	DeleteReviewByOwner(ctx context.Context, userID int64, offerID int64) error
	GetReview(ctx context.Context, reviewID int64) (Review, error)
	GetReviewByOwner(ctx context.Context, userID int64, offerID int64) (Review, error)
	GetReviewsOfOffer(ctx context.Context, offerID int64) ([]*Review, error)
	IsReviewsBelongToUser(ctx context.Context, reviewID int64, userID int64) (bool, error)
	IsReviewsBelongToOffer(ctx context.Context, reviewID int64, offerID int64) (bool, error)
}

// ReviewService is a service which is responsible for actions related to review
type ReviewService interface{}
