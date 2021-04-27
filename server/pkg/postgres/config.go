package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Config contains required values for initializing pgxpool.Config. Use in newConnConfig.
type Config struct {
	Host     string `required:"true"`
	Port     uint16 `required:"true"`
	Database string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
}

// NewPool initializes postgres pool. If error is returned, cleanup is nil.
func NewPool(conf Config) (*pgxpool.Pool, func(), error) {
	connURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.User, conf.Password),
		Host:   fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Path:   conf.Database,
	}

	poolConf, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not initialize pool config: %w", err)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), poolConf)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: could not initialize pool: %w", err)
	}

	return pool, pool.Close, nil
}
