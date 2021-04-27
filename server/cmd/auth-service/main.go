package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pw-software-engineering/b-team/server/pkg/app"
	"github.com/pw-software-engineering/b-team/server/pkg/postgres"
	"go.uber.org/zap"
)

type config struct {
	App             app.Config
	Postgres        postgres.Config
	SessionDuration time.Duration `envconfig:"SESSION_DURATION" default:"24h"`
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panicf("could not initialize logger: %s", err.Error())
	}
	var cfg config
	app.LoadConfig(logger, &cfg)

	application := app.NewApp(logger, cfg.App)
	application.Build(func() error {
		pool, cleanup, err := postgres.NewPool(cfg.Postgres)
		if err != nil {
			return fmt.Errorf("could not initialize postgres connection pool: %w", err)
		}
		application.AddCleanup(cleanup)
		application.AddHealthCheck(postgres.NewHealthConfig(pool))

		sessionStorage := postgres.NewSessionStorage(pool, cfg.SessionDuration)

		api := newAPI(application.Logger, NewSessionVerifier(sessionStorage))
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
