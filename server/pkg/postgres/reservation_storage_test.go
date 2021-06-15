package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CleanTestReservationStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM reservations",
		"DELETE FROM room_reservations",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestReservationStorage_CreateReservation(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	reservation := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    1,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   5,
	}

	_, err := storage.CreateReservation(ctx, reservation)
	require.NoError(t, err)
}

func TestReservationStorage_DeleteReservation(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	reservation := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    1,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   5,
	}

	id, errAdd := storage.CreateReservation(ctx, reservation)
	require.NoError(t, errAdd)
	errDelete := storage.DeleteReservation(ctx, id)
	require.NoError(t, errDelete)
}

func TestReservationStorage_GetClientReservations(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	reservation0 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    1,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   1,
	}
	reservation1 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    2,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   1,
	}
	reservation2 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    2,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   2,
	}

	_, errAdd0 := storage.CreateReservation(ctx, reservation0)
	require.NoError(t, errAdd0)
	_, errAdd1 := storage.CreateReservation(ctx, reservation1)
	require.NoError(t, errAdd1)
	id2, errAdd2 := storage.CreateReservation(ctx, reservation2)
	require.NoError(t, errAdd2)

	emptyList, errNo := storage.GetClientReservations(ctx, 256)
	require.NoError(t, errNo)
	assert.Equal(t, 0, len(emptyList))
	list2, err := storage.GetClientReservations(ctx, 2)
	require.NoError(t, err)
	require.Equal(t, 1, len(list2))
	assert.Equal(t, id2, list2[0].ID)
}

func TestReservationStorage_GetHotelReservations(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	reservation0 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    1,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   1,
	}
	reservation1 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    2,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   1,
	}
	reservation2 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    2,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   2,
	}

	id0, errAdd0 := storage.CreateReservation(ctx, reservation0)
	require.NoError(t, errAdd0)
	_, errAdd1 := storage.CreateReservation(ctx, reservation1)
	require.NoError(t, errAdd1)
	_, errAdd2 := storage.CreateReservation(ctx, reservation2)
	require.NoError(t, errAdd2)

	emptyList, errNo := storage.GetHotelReservations(ctx, 256)
	require.NoError(t, errNo)
	assert.Equal(t, 0, len(emptyList))
	list1, err := storage.GetHotelReservations(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 1, len(list1))
	assert.Equal(t, id0, list1[0].ID)
}

func TestReservationStorage_IsReservationOwnedByClient(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	reservation0 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    1,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   1,
	}
	reservation1 := &bookly.Reservation{
		FromTime:   time.Now(),
		ToTime:     time.Now(),
		HotelID:    2,
		OfferID:    1,
		ChildCount: 5,
		AdultCount: 2,
		ClientID:   5,
	}

	id0, errAdd0 := storage.CreateReservation(ctx, reservation0)
	require.NoError(t, errAdd0)
	id1, errAdd1 := storage.CreateReservation(ctx, reservation1)
	require.NoError(t, errAdd1)

	_, errNotExists := storage.IsReservationOwnedByClient(ctx, 1, 142134)
	require.Error(t, errNotExists)
	owned1, errNotOwned := storage.IsReservationOwnedByClient(ctx, 1, id1)
	require.NoError(t, errNotOwned)
	assert.Equal(t, false, owned1)
	owned2, err := storage.IsReservationOwnedByClient(ctx, 1, id0)
	require.NoError(t, err)
	assert.Equal(t, true, owned2)
}

func TestReservationStorage_IsRoomBooked(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewReservationStorage(initDb(t))
	ctx := context.Background()
	CleanTestReservationStorage(t, storage.connPool, ctx)

	errAdd1 := storage.CreateReservationRoomLink(ctx, 1, 15)
	require.NoError(t, errAdd1)

	notBooked, err1 := storage.IsRoomBooked(ctx, 50)
	require.NoError(t, err1)
	assert.False(t, notBooked)
	booked, err2 := storage.IsRoomBooked(ctx, 15)
	require.NoError(t, err2)
	assert.True(t, booked)
}
