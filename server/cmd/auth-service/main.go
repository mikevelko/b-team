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
		storage, cleanup, err := postgres.NewSessionStorage(cfg.Postgres, cfg.SessionDuration)
		if err != nil {
			return fmt.Errorf("could not initialize postgres: %w", err)
		}
		application.AddCleanup(cleanup)
		api := newAPI(application.Logger, NewSessionVerifier(storage))
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
