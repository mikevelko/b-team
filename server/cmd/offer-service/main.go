package main

import (
	"fmt"
	"log"

	"github.com/pw-software-engineering/b-team/server/pkg/app"
	"github.com/pw-software-engineering/b-team/server/pkg/postgres"
	"go.uber.org/zap"
)

var cfg struct {
	App      app.Config
	Postgres postgres.Config
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panicf("could not initialize logger: %s", err.Error())
	}

	application := app.NewApp(logger, cfg.App)
	application.Build(func() error {
		// we do dependency injection here
		storage, cleanup, err := postgres.NewOfferStorage(cfg.Postgres)
		if err != nil {
			return fmt.Errorf("could not initialize postgres: %w", err)
		}
		defer cleanup() // todo: add cleanup handling to app
		service := newOfferService(storage)
		api := newAPI(application.Logger, service)
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
