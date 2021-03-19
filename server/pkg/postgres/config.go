package postgres

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4/pgxpool"
    "net/url"
)

// Config contains required values for initializing pgxpool.Config. Use in newConnConfig.
type Config struct {
    Host     string
    Port     uint16
    Database string
    User     string
    Password string
}

// newPool initializes postgres pool. If error is returned, cleanup is nil.
func newPool(conf Config) (*pgxpool.Pool, func(), error) {
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
