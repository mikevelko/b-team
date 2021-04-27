package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func CleanTestOfferStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM offers",
		"DELETE FROM rooms",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestOfferStorage_CreateGetDeleteOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))
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

func TestOfferStorage_GetAllOffers(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))
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
		IsActive:            false,
		OfferTitle:          "2 Rooms With Private Beach",
		CostPerChild:        decimal.New(1234, -2),
		CostPerAdult:        decimal.New(1234, -2),
		MaxGuests:           5,
		Description:         "Very good offer from our hotel",
		OfferPreviewPicture: "https:__9gag.com_gag_aEpgjL9", // todo: change it, when support for pictures is added
		Pictures:            nil,                            // todo: change it, when support for pictures is added
		Rooms:               nil,                            // todo: change it when support for rooms is added
	})
	hotelLinks := [3]int64{1, 2, 1}
	ctx := context.Background()
	CleanTestOfferStorage(t, storage.connPool, ctx)

	for i, o := range offers {
		id, errAdd := storage.CreateOffer(ctx, o, hotelLinks[i])
		offers[i].ID = id
		require.NoError(t, errAdd)
	}
	resultAll, errGetAll := storage.GetAllOffers(ctx, 1, nil)
	require.NoError(t, errGetAll)
	assert.Equal(t, 2, len(resultAll))
	for i, o := range offers {
		if hotelLinks[i] == 1 {
			assert.Contains(t, resultAll, o)
		}
	}
	onlyActive := true
	resultActive, errGetActive := storage.GetAllOffers(ctx, 1, &onlyActive)
	require.NoError(t, errGetActive)
	assert.Equal(t, 1, len(resultActive))
	for i, o := range offers {
		if hotelLinks[i] == 1 && offers[i].IsActive == onlyActive {
			assert.Contains(t, resultActive, o)
		}
	}
	onlyInactive := true
	resultInactive, errGetInactive := storage.GetAllOffers(ctx, 1, &onlyInactive)
	require.NoError(t, errGetInactive)
	assert.Equal(t, 1, len(resultInactive))
	for i, o := range offers {
		if hotelLinks[i] == 1 && offers[i].IsActive == onlyInactive {
			assert.Contains(t, resultInactive, o)
		}
	}
}

func TestOfferStorage_GetSpecificOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))

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
	CleanTestOfferStorage(t, storage.connPool, ctx)
	offerID, errAdd := storage.CreateOffer(ctx, correctOffer, 1)
	require.NoError(t, errAdd)
	_, errNotExists := storage.GetSpecificOffer(ctx, 10000)
	require.Error(t, errNotExists)
	offer, errRetrieve := storage.GetSpecificOffer(ctx, offerID)
	require.NoError(t, errRetrieve)
	assert.Equal(t, correctOffer.OfferTitle, offer.OfferTitle)
}

func TestOfferStorage_UpdateOfferDetails(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))

	preOffer := &bookly.Offer{
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
	newOffer := &bookly.Offer{
		IsActive:            true,
		OfferTitle:          "NewSometitle",
		CostPerChild:        decimal.New(1234, -2), // it its 12.32
		CostPerAdult:        decimal.New(4321, -2),
		MaxGuests:           2,
		Description:         "dfdfsdfsd",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	}
	ctx := context.Background()
	CleanTestOfferStorage(t, storage.connPool, ctx)
	offerID, errAdd := storage.CreateOffer(ctx, preOffer, 1)
	require.NoError(t, errAdd)
	errUpdate := storage.UpdateOfferDetails(ctx, offerID, *newOffer)
	require.NoError(t, errUpdate)
	offer, errRetrieve := storage.GetSpecificOffer(ctx, offerID)
	require.NoError(t, errRetrieve)
	assert.Equal(t, newOffer.OfferTitle, offer.OfferTitle)
}

func TestOfferStorage_IsOfferActive(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))

	offerActive := &bookly.Offer{
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
	offerInactive := &bookly.Offer{
		IsActive:            false,
		OfferTitle:          "NewSometitle",
		CostPerChild:        decimal.New(1234, -2), // it its 12.32
		CostPerAdult:        decimal.New(4321, -2),
		MaxGuests:           2,
		Description:         "dfdfsdfsd",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	}
	ctx := context.Background()
	CleanTestOfferStorage(t, storage.connPool, ctx)
	activeID, errAddActive := storage.CreateOffer(ctx, offerActive, 1)
	require.NoError(t, errAddActive)
	inactiveID, errAddInactive := storage.CreateOffer(ctx, offerInactive, 1)
	require.NoError(t, errAddInactive)
	active1, err1 := storage.IsOfferActive(ctx, activeID)
	require.NoError(t, err1)
	assert.Equal(t, true, active1)
	active2, err2 := storage.IsOfferActive(ctx, inactiveID)
	require.NoError(t, err2)
	assert.Equal(t, false, active2)
	_, errNotFound := storage.IsOfferActive(ctx, 14141231)
	require.Error(t, errNotFound)
}

func TestOfferStorage_IsOfferOwnedByHotel(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))

	offerActive := &bookly.Offer{
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
	offerInactive := &bookly.Offer{
		IsActive:            false,
		OfferTitle:          "NewSometitle",
		CostPerChild:        decimal.New(1234, -2), // it its 12.32
		CostPerAdult:        decimal.New(4321, -2),
		MaxGuests:           2,
		Description:         "dfdfsdfsd",
		OfferPreviewPicture: "http://localhost:/mypicture123", // todo: change it, when support for pictures is added
		Pictures:            nil,                              // todo: change it, when support for pictures is added
		Rooms:               nil,                              // todo: change it when support for rooms is added
	}
	ctx := context.Background()
	CleanTestOfferStorage(t, storage.connPool, ctx)
	ownedID, errAddActive := storage.CreateOffer(ctx, offerActive, 1)
	require.NoError(t, errAddActive)
	notOwnedID, errAddInactive := storage.CreateOffer(ctx, offerInactive, 44)
	require.NoError(t, errAddInactive)

	owned1, err1 := storage.IsOfferOwnedByHotel(ctx, 1, ownedID)
	require.NoError(t, err1)
	assert.Equal(t, true, owned1)
	owned2, err2 := storage.IsOfferOwnedByHotel(ctx, 1, notOwnedID)
	require.NoError(t, err2)
	assert.Equal(t, false, owned2)
	_, errNotFound := storage.IsOfferOwnedByHotel(ctx, 1, 14141231)
	require.Error(t, errNotFound)
}

func TestOfferStorage_MarkDeletedOfferAndCheckStatus(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewOfferStorage(initDb(t))

	offerActive := &bookly.Offer{
		IsActive:            false,
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
	CleanTestOfferStorage(t, storage.connPool, ctx)
	ID, errAdd := storage.CreateOffer(ctx, offerActive, 1)
	require.NoError(t, errAdd)

	errMark := storage.SetOfferDeletionStatus(ctx, ID, true)
	require.NoError(t, errMark)
	deleted, errCheck := storage.IsOfferMarkedAsDeleted(ctx, ID)
	require.NoError(t, errCheck)
	assert.Equal(t, true, deleted)
}
