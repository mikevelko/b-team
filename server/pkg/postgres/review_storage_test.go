package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

func CleanTestReviewStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM reviews",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestReviewStorage_CreateReview(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Time{},
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	_ = reviewID
}

func TestReviewStorage_DeleteReview(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Time{},
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	_ = reviewID

	err = storage.DeleteReview(ctx, reviewID+1)
	require.Error(t, err)
	assert.ErrorIs(t, bookly.ErrReviewNotFound, err)

	err = storage.DeleteReview(ctx, reviewID)
	require.NoError(t, err)
}

func TestReviewStorage_DeleteReviewByOwner(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Time{},
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	_ = reviewID

	err = storage.DeleteReviewByOwner(ctx, 1, 1)
	require.Error(t, err)
	assert.ErrorIs(t, bookly.ErrReviewNotFound, err)

	err = storage.DeleteReviewByOwner(ctx, 0, 0)
	require.NoError(t, err)
}

func TestReviewStorage_GetReview(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Now(),
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	review.ID = reviewID

	_, err = storage.GetReview(ctx, reviewID+1)
	require.Error(t, err)
	assert.ErrorIs(t, bookly.ErrReviewNotFound, err)

	/*getReview*/
	_, err = storage.GetReview(ctx, reviewID)
	require.NoError(t, err)
	// require.Equal(t, review, getReview)
	// todo: Add correct time
}

func TestReviewStorage_GetReviewByOwner(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Now(),
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	review.ID = reviewID

	_, err = storage.GetReviewByOwner(ctx, 1, 1)
	require.Error(t, err)
	assert.ErrorIs(t, bookly.ErrReviewNotFound, err)

	/*getReview*/
	_, err = storage.GetReviewByOwner(ctx, 0, 0)
	require.NoError(t, err)
	// require.Equal(t, review, getReview)
	// todo: Add correct time
}

func TestReviewStorage_GetReviewsOfOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReviewStorage(initDb(t))

	ctx := context.Background()
	CleanTestReviewStorage(t, storage.connPool, ctx)

	review := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     0,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Now(),
	}
	reviewID, err := storage.CreateReview(ctx, review)
	require.NoError(t, err)
	review.ID = reviewID

	review2 := bookly.Review{
		ID:         0,
		OfferID:    0,
		UserID:     1,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Now(),
	}
	reviewID, err = storage.CreateReview(ctx, review2)
	require.NoError(t, err)
	review2.ID = reviewID

	review3 := bookly.Review{
		ID:         0,
		OfferID:    1,
		UserID:     1,
		Content:    "",
		Rating:     0,
		ReviewDate: time.Now(),
	}
	reviewID, err = storage.CreateReview(ctx, review3)
	require.NoError(t, err)
	review3.ID = reviewID

	reviews, err := storage.GetReviewsOfOffer(ctx, 0)
	require.NoError(t, err)
	require.Equal(t, 2, len(reviews))
}
