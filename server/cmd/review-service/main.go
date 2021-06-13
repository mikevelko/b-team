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

		reviewStorage := postgres.NewReviewStorage(pool)
		userStorage := postgres.NewUserStorage(pool)
		offerStorage := postgres.NewOfferStorage(pool)
		reservationStorage := postgres.NewReservationStorage(pool)

		service := newReviewService(reviewStorage, offerStorage, userStorage, reservationStorage)

		api := newAPI(application.Logger, service)
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
