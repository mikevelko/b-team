package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// ReviewStorage is responsible for storing and retrieving offers
type ReviewStorage struct {
	connPool *pgxpool.Pool
}

// this is very nice way of ensuring, that *ReviewStorage{} implements bookly.ReviewStorage.
// if it does not - program won't compile and you'll see red error in IDE
var _ bookly.ReviewStorage = &ReviewStorage{}

// NewReviewStorage initializes ReviewStorage
func NewReviewStorage(pool *pgxpool.Pool) *ReviewStorage {
	return &ReviewStorage{
		connPool: pool,
	}
}

// CreateReview implements database feat of creating reviews
func (r *ReviewStorage) CreateReview(ctx context.Context, review bookly.Review) (int64, error) {
	const query = `
    INSERT INTO reviews(
        user_id,
    	offer_id,
    	content,
    	rating,
    	review_date
        )
    VALUES ($1,$2,$3,$4,$5)
    RETURNING id;
`
	var id int64
	err := r.connPool.QueryRow(ctx, query,
		review.UserID,
		review.OfferID,
		review.Content,
		review.Rating,
		review.ReviewDate,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: could not insert review: %w", err)
	}
	return id, nil
}

// DeleteReview implements database feat of deleting reviews
func (r *ReviewStorage) DeleteReview(ctx context.Context, reviewID int64) error {
	const queryDelete = `
    DELETE
	FROM reviews
	WHERE id = $1
`
	log.Println("database func")
	rows, err := r.connPool.Exec(ctx, queryDelete,
		reviewID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.ErrReviewNotFound
		}
		return fmt.Errorf("postgres: could not delete record from reviews: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return bookly.ErrReviewNotFound
	}

	return nil
}

// DeleteReviewByOwner implements database feat of deleting reviews
func (r *ReviewStorage) DeleteReviewByOwner(ctx context.Context, userID int64, offerID int64) error {
	const queryDelete = `
    DELETE
	FROM reviews
	WHERE user_id = $1 AND offer_id = $2
`
	log.Println("database func")
	rows, err := r.connPool.Exec(ctx, queryDelete,
		userID,
		offerID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.ErrReviewNotFound
		}
		return fmt.Errorf("postgres: could not delete record from reviews: %w", err)
	}
	if rows.RowsAffected() == 0 {
		return bookly.ErrReviewNotFound
	}

	return nil
}

// GetReview implements database feat of getting reviews
func (r *ReviewStorage) GetReview(ctx context.Context, reviewID int64) (bookly.Review, error) {
	const queryCheck = `
    SELECT *
	FROM reviews
	WHERE id = $1
`

	row := r.connPool.QueryRow(ctx, queryCheck,
		reviewID,
	)
	review := bookly.Review{}
	err := row.Scan(&review.ID, &review.UserID, &review.OfferID, &review.Content, &review.Rating, &review.ReviewDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.Review{}, bookly.ErrReviewNotFound
		}
		return bookly.Review{}, fmt.Errorf("postgres: could not get review: %w", err)
	}

	return review, nil
}

// GetReviewByOwner implements database feat of getting reviews
func (r *ReviewStorage) GetReviewByOwner(ctx context.Context, userID int64, offerID int64) (bookly.Review, error) {
	const queryCheck = `
    SELECT *
	FROM reviews
	WHERE user_id = $1 AND offer_id = $2
`

	row := r.connPool.QueryRow(ctx, queryCheck,
		userID,
		offerID,
	)
	review := bookly.Review{}
	err := row.Scan(&review.ID, &review.UserID, &review.OfferID, &review.Content, &review.Rating, &review.ReviewDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return bookly.Review{}, bookly.ErrReviewNotFound
		}
		return bookly.Review{}, fmt.Errorf("postgres: could not get review: %w", err)
	}

	return review, nil
}

// GetReviewsOfOffer implements database feat of getting reviews
func (r *ReviewStorage) GetReviewsOfOffer(ctx context.Context, offerID int64) ([]*bookly.Review, error) {
	const query = `
    SELECT *
	FROM reviews
	WHERE offer_id = $1
`

	list, err := r.connPool.Query(ctx, query,
		offerID,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: could not get reviews: %w", err)
	}

	result := []*bookly.Review{}
	defer list.Close()
	for list.Next() {
		var review bookly.Review
		err = list.Scan(&review.ID, &review.UserID, &review.OfferID, &review.Content, &review.Rating, &review.ReviewDate)
		result = append(result, &review)
	}
	errFinal := list.Err()
	if errFinal != nil {
		return nil, fmt.Errorf("postgres: could not get reviews: %w", err)
	}

	return result, nil
}

// IsReviewsBelongToUser implements database feat of checking reviews
func (r *ReviewStorage) IsReviewsBelongToUser(ctx context.Context, reviewID int64, userID int64) (bool, error) {
	review, err := r.GetReview(ctx, reviewID)
	if err != nil {
		return false, err
	}
	return review.UserID == userID, nil
}

// IsReviewsBelongToOffer implements database feat of checking reviews
func (r *ReviewStorage) IsReviewsBelongToOffer(ctx context.Context, reviewID int64, offerID int64) (bool, error) {
	review, err := r.GetReview(ctx, reviewID)
	if err != nil {
		return false, err
	}
	return review.OfferID == offerID, nil
}
