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
		// we do dependency injection here
		/*roomStorage*/
		roomStorage, cleanupUser, errU := postgres.NewRoomStorage(cfg.Postgres)
		if errU != nil {
			return fmt.Errorf("could not initialize postgres: %w", err)
		}
		application.AddCleanup(cleanupUser)

		service := newRoomService(roomStorage)
		api := newAPI(application.Logger, service)
		api.mount(application.Router)
		return nil
	})
	application.Run()
}
