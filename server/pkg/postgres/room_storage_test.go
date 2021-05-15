package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

func CleanTestRoomStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM rooms",
		"DELETE FROM offers_rooms",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestRoomStorage_CreateRoom(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	correctRoom := bookly.Room{
		ID:         0,
		RoomNumber: "12F",
		HotelID:    0,
	}
	roomID, err := storage.CreateRoom(ctx, correctRoom, 1)
	require.NoError(t, err)
	_ = roomID

	incorrectRoom := bookly.Room{
		ID:         0,
		RoomNumber: "12F",
		HotelID:    0,
	}
	roomID, err = storage.CreateRoom(ctx, incorrectRoom, 1)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomAlreadyExists)
}

func TestRoomStorage_DeleteRoom(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))
	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	correctRoom := bookly.Room{
		ID:         0,
		RoomNumber: "12Fa",
		HotelID:    0,
	}
	_, err := storage.CreateRoom(ctx, correctRoom, 1)
	require.NoError(t, err)
	correctRoom = bookly.Room{
		ID:         0,
		RoomNumber: "13Fa",
		HotelID:    0,
	}
	room1, err := storage.CreateRoom(ctx, correctRoom, 1)
	require.NoError(t, err)
	correctRoom = bookly.Room{
		ID:         0,
		RoomNumber: "14Fa",
		HotelID:    0,
	}
	room2, err := storage.CreateRoom(ctx, correctRoom, 2)
	require.NoError(t, err)

	err = storage.DeleteRoom(ctx, 123, 1)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomNotFound)

	err = storage.DeleteRoom(ctx, room2, 1)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomNotBelongToHotel)

	err = storage.DeleteRoom(ctx, room1, 1)
	require.NoError(t, err)

	return
}

func TestRoomStorage_GetAllHotelRoom(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))
	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	correctRoom1 := bookly.Room{
		ID:         0,
		RoomNumber: "12Fa",
		HotelID:    0,
	}
	_, err := storage.CreateRoom(ctx, correctRoom1, 1)
	require.NoError(t, err)
	correctRoom2 := bookly.Room{
		ID:         0,
		RoomNumber: "13Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom2, 1)
	require.NoError(t, err)
	correctRoom3 := bookly.Room{
		ID:         0,
		RoomNumber: "14Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom3, 2)
	require.NoError(t, err)

	_, err = storage.GetAllHotelRooms(ctx, 1)
	require.NoError(t, err)
	return
}

func TestRoomStorage_GetRoomByName(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	correctRoom1 := bookly.Room{
		ID:         0,
		RoomNumber: "12Fa",
		HotelID:    0,
	}
	_, err := storage.CreateRoom(ctx, correctRoom1, 1)
	require.NoError(t, err)
	correctRoom2 := bookly.Room{
		ID:         0,
		RoomNumber: "13Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom2, 1)
	require.NoError(t, err)
	correctRoom3 := bookly.Room{
		ID:         0,
		RoomNumber: "14Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom3, 2)
	require.NoError(t, err)

	_, err = storage.GetRoomByName(ctx, "14Fa", 1)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomNotFound)

	_, err = storage.GetRoomByName(ctx, "14Fa", 2)
	require.NoError(t, err)

	return
}

func TestRoomStorage_GetRoom(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	correctRoom1 := bookly.Room{
		ID:         0,
		RoomNumber: "12Fa",
		HotelID:    0,
	}
	id, err := storage.CreateRoom(ctx, correctRoom1, 1)
	require.NoError(t, err)
	correctRoom2 := bookly.Room{
		ID:         0,
		RoomNumber: "13Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom2, 1)
	require.NoError(t, err)
	correctRoom3 := bookly.Room{
		ID:         0,
		RoomNumber: "14Fa",
		HotelID:    0,
	}
	_, err = storage.CreateRoom(ctx, correctRoom3, 2)
	require.NoError(t, err)

	_, err = storage.GetRoom(ctx, 500)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomNotFound)

	_, err = storage.GetRoom(ctx, id)
	require.NoError(t, err)

	return
}

func TestRoomStorage_AddLinkWithRoomAndOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	err := storage.AddLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)

	return
}

func TestRoomStorage_DeleteLinkWithRoomAndOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	err := storage.AddLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)

	err = storage.DeleteLinkWithRoomAndOffer(ctx, 2, 1)
	require.Error(t, err)
	assert.ErrorIs(t, bookly.ErrLinkOfferRoomNotFound, err)

	err = storage.DeleteLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)

	return
}

func TestRoomStorage_IsExistLinkWithRoomAndOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	err := storage.AddLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)

	res, err := storage.IsExistLinkWithRoomAndOffer(ctx, 2, 1)
	require.NoError(t, err)
	assert.Equal(t, false, res)

	res, err = storage.IsExistLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, true, res)

	return
}

func TestRoomStorage_OffersRelatedWithRoom(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	err := storage.AddLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 2, 1)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 3, 1)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 4, 1)
	require.NoError(t, err)

	list, err := storage.OffersRelatedWithRoom(ctx, 2)
	require.NoError(t, err)
	assert.Equal(t, 0, len(list))

	list, err = storage.OffersRelatedWithRoom(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, 4, len(list))

	return
}

func TestRoomStorage_RoomsRelatedWithOffer(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewRoomStorage(initDb(t))

	ctx := context.Background()
	CleanTestRoomStorage(t, storage.connPool, ctx)

	err := storage.AddLinkWithRoomAndOffer(ctx, 1, 1)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 1, 2)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 1, 3)
	require.NoError(t, err)
	err = storage.AddLinkWithRoomAndOffer(ctx, 1, 4)
	require.NoError(t, err)

	list, err := storage.RoomsRelatedWithRoom(ctx, 2)
	require.NoError(t, err)
	assert.Equal(t, 0, len(list))

	list, err = storage.RoomsRelatedWithRoom(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, 4, len(list))

	return
}
