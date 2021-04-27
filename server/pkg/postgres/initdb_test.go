package postgres

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

// conf is used as a default when running integration tests with docker-compose
var conf Config = Config{
	Host:     "localhost",
	Port:     5432,
	Database: "bookly",
	User:     "bookly",
	Password: "bookly",
}

func initDb(t *testing.T) *pgxpool.Pool {
	db, cleanup, err := NewPool(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	return db
}
