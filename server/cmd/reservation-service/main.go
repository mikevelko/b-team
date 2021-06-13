package main

import (
	"fmt"
	"log"

	"github.com/pw-software-engineering/b-team/server/pkg/app"
	"github.com/pw-software-engineering/b-team/server/pkg/postgres"
	"go.uber.org/zap"
)

type config struct {
	App      app.Config
	Postgres postgres.Config
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

		offerStorage := postgres.NewOfferStorage(pool)
		reservationStorage := postgres.NewReservationStorage(pool)
		hotelStorage := postgres.NewHotelStorage(pool)
		reviewStorage := postgres.NewReviewStorage(pool)

		service := newReservationService(offerStorage, reservationStorage, hotelStorage, reviewStorage)
		api := newAPI(application.Logger, service)
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
