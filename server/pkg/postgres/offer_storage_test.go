package postgres

import (
	"context"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/rently"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestOfferStorage_CreateGetDeleteOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewOfferStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	correctOffer := &rently.Offer{
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
	offerID, err := storage.CreateOffer(ctx, correctOffer)
	require.NoError(t, err)

	// todo: when support for getting offers is added, get it from storage
	_ = offerID
}

func TestOfferStorage_AddUpdateOfferStatus(t *testing.T) {
	testutils.SetIntegration(t)
	storage, cleanup, err := NewOfferStorage(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	correctOffer := &rently.Offer{
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
	offerID, errAdd := storage.CreateOffer(ctx, correctOffer)
	require.NoError(t, errAdd)
	errUpdate := storage.UpdateOfferStatus(ctx, offerID, false)
	require.NoError(t, errUpdate)
	// todo: when getting offers is implemented, get changed offer and assert that IsActive is false
	_ = offerID
}
