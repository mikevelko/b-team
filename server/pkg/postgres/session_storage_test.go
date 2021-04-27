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

func CleanSessionTestStorage(t *testing.T, pool *pgxpool.Pool, ctx context.Context) {
	queries := []string{
		"DELETE FROM sessions",
		"DELETE FROM users",
	}
	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		require.NoError(t, err)
	}
}

func TestSessionStorage_CreateNew(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewSessionStorage(initDb(t), time.Duration(10)*time.Minute)

	ctx := context.Background()
	CleanSessionTestStorage(t, storage.connPool, ctx)
	token, err := storage.CreateNew(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, (int64)(1), token.ID)
}

func TestSessionStorage_Update(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewSessionStorage(initDb(t), time.Duration(10)*time.Minute)
	ctx := context.Background()
	CleanSessionTestStorage(t, storage.connPool, ctx)
	const forceAddQuery = `
    INSERT INTO sessions(
	creation_date,
	expire_date,
	user_id)
	VALUES($1,$2,$3)
`
	tokenUp, errAddNow := storage.CreateNew(ctx, 10)
	require.NoError(t, errAddNow)
	expireTime := time.Now().Local().Add(-20 * time.Minute).Format(time.RFC3339)
	tokenExp := bookly.Token{ID: 11, CreatedAt: time.Now().Format(time.RFC3339)}
	_, errAddExp := storage.connPool.Exec(ctx, forceAddQuery, tokenExp.CreatedAt, expireTime, 11)

	require.NoError(t, errAddExp)
	errUpNow := storage.Update(ctx, tokenUp)
	assert.NoError(t, errUpNow)
	errUpExp := storage.Update(ctx, tokenExp)
	assert.Equal(t, bookly.ErrSessionExpired, errUpExp)
	errNotEx := storage.Update(ctx, bookly.Token{ID: 50, CreatedAt: "Asdsadsad-123021"})
	assert.Equal(t, bookly.ErrSessionNotFound, errNotEx)
}

func TestSessionStorage_GetSession(t *testing.T) {
	testutils.SetIntegration(t)
	storage := NewSessionStorage(initDb(t), time.Duration(2000)*time.Minute)

	ctx := context.Background()
	CleanSessionTestStorage(t, storage.connPool, ctx)
	const forceAddQuery = `
    INSERT INTO users(
	first_name,
	surname,
	email,
	user_name,
	password,
	hotel_id)
	VALUES($1,$2,$3,$4,$5,$6)
	RETURNING id;
`
	var janID int64
	var traczID int64
	errNoHotel := storage.connPool.QueryRow(ctx, forceAddQuery, "Jan", "Kowalski",
		"jan.kowal@rmail.xd", "Janek777", "1253", nil).Scan(&janID)
	require.NoError(t, errNoHotel)
	errHotel := storage.connPool.QueryRow(ctx, forceAddQuery, "Janusz",
		"Tracz", "tracz.boss@rmail.xd", "Trancjusz", "Plebania1829", 1).Scan(&traczID)
	require.NoError(t, errHotel)

	hotelIdNil, errGetNil := storage.GetSession(ctx, bookly.Token{ID: janID, CreatedAt: "XDDDD"})
	require.NoError(t, errGetNil)
	assert.Equal(t, int64(0), hotelIdNil.HotelID)
	hotelId, errGet := storage.GetSession(ctx, bookly.Token{ID: traczID, CreatedAt: "ASDASDSA"})
	require.NoError(t, errGet)
	assert.Equal(t, int64(1), hotelId.HotelID)
}
