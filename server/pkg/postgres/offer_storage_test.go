package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func CleanTestStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM offers",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestOfferStorage_CreateGetDeleteOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewOfferStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	correctOffer := &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "Sometitle",
		CostPerChild:        decimal.New(1234, -2), // it its 12.32
		CostPerAdult:        decimal.New(4321, -2),
		MaxGuests:           2,
		Description:         "dfdfsdfsd",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	}
	ctx := context.Background()
	offerID, err := storage.CreateOffer(ctx, correctOffer, 1)
	require.NoError(t, err)

	// todo: when support for getting offers is added, get it from storage
	_ = offerID
}

func TestOfferStorage_AddUpdateOfferStatus(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewOfferStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	correctOffer := &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "Sometitle",
		CostPerChild:        decimal.New(1234, -2), // it its 12.32
		CostPerAdult:        decimal.New(4321, -2),
		MaxGuests:           2,
		Description:         "dfdfsdfsd",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	}
	ctx := context.Background()
	CleanTestStorage(t, storage.connPool, ctx)
	offerID, errAdd := storage.CreateOffer(ctx, correctOffer, 1)
	require.NoError(t, errAdd)
	errUpdate := storage.UpdateOfferStatus(ctx, offerID, false)
	require.NoError(t, errUpdate)
	// todo: when getting offers is implemented, get changed offer and assert that IsActive is false
	_ = offerID
}

func TestOfferStorage_GetAllOffers(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewOfferStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	var offers []*bookly.Offer
	offers = append(offers, &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "2 Room Comfort",
		CostPerChild:        decimal.New(1234, -2),
		CostPerAdult:        decimal.New(1234, -2),
		MaxGuests:           2,
		Description:         "Very good offer from our hotel",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	})
	offers = append(offers, &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "4 Room Luxury",
		CostPerChild:        decimal.New(1234, -2),
		CostPerAdult:        decimal.New(1234, -2),
		MaxGuests:           4,
		Description:         "Very good offer from our mortal ene... competition",
		OfferPreviewPicture: "http://localhost:/mypicture2137", // todo: change it, when support for pictures is added
		Pictures:            nil,                               // todo: change it, when support for pictures is added
		Rooms:               nil,                               // todo: change it when support for rooms is added
	})
	offers = append(offers, &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "2 Rooms With Private Beach",
		CostPerChild:        decimal.New(1234, -2),
		CostPerAdult:        decimal.New(1234, -2),
		MaxGuests:           5,
		Description:         "Very good offer from our hotel",
		OfferPreviewPicture: "https:__9gag.com_gag_aEpgjL9", // todo: change it, when support for pictures is added
		Pictures:            nil,                            // todo: change it, when support for pictures is added
		Rooms:               nil,                            // todo: change it when support for rooms is added
	})
	hotelLinks := [3]int{1, 2, 1}
	ctx := context.Background()
	CleanTestStorage(t, storage.connPool, ctx)

	for i, o := range offers {
		_, errAdd := storage.CreateOffer(ctx, o, hotelLinks[i])
		require.NoError(t, errAdd)
	}
	result, errGet := storage.GetAllOffers(ctx, 1)
	require.NoError(t, errGet)
	assert.Equal(t, 2, len(result))
	for i, o := range offers {
		if hotelLinks[i] == 1 {
			assert.Contains(t, result, o)
		}
	}
}
