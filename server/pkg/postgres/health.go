package postgres

import (
	"context"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewHealthConfig initializes configuration for postgres health check
func NewHealthConfig(pool *pgxpool.Pool) *gosundheit.Config {
	return &gosundheit.Config{
		Check:            newCheck(pool),
		ExecutionPeriod:  30 * time.Second,
		InitialDelay:     0,
		InitiallyPassing: false,
	}
}

type healthCheck struct {
	pool *pgxpool.Pool
}

func newCheck(pool *pgxpool.Pool) *healthCheck {
	return &healthCheck{pool: pool}
}

var _ checks.Check = &healthCheck{}

func (c *healthCheck) Name() string {
	return "postgres"
}

func (c *healthCheck) Execute() (details interface{}, err error) {
	details = c.pool.Config().ConnString()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = c.pool.Exec(ctx, `select 1;`)
	return details, err
}
