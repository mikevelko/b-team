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

	_, err = storage.GetRoom(ctx, "14Fa", 1)
	require.Error(t, err)
	assert.ErrorIs(t, err, bookly.ErrRoomNotFound)

	_, err = storage.GetRoom(ctx, "14Fa", 2)
	require.NoError(t, err)

	return
}
