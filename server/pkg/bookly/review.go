package bookly

import (
	"context"
	"time"
)

// Review is a model for room, not implemented yed
type Review struct {
	ID         int64     `json:"reviewID"`
	OfferID    int64     `json:"-"`
	UserID     int64     `json:"-"`
	Content    string    `json:"content"`
	Rating     int32     `json:"rating"` //[1-5]
	ReviewDate time.Time `json:"creationDate"`

	ReviewerUsername string `json:"reviewerUsername"` // put in service
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
type ReviewService interface {
	GetReviewsOfOffer(ctx context.Context, offerID int64, userID int64) ([]*Review, error)
	CreateReview(ctx context.Context, review Review, userID int64, offerID int64) (int64, error)
	DeleteReview(ctx context.Context, reviewID int64, userID int64, offerID int64) error
	UpdateReview(ctx context.Context, review Review, userID int64, offerID int64) (int64, error)
}
