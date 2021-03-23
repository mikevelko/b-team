package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

//Config is used to configure application
type Config struct {
	HTTPHostAddress string `env:"default=0.0.0.0:8080"`
}

//NewApp initializes application
func NewApp(logger *zap.Logger, cfg Config) *App {
	return &App{
		Logger: logger,
		Router: chi.NewRouter(),
		config: cfg,
	}
}

// App is responsible for initializing and running application
type App struct {
	Logger     *zap.Logger
	Router     chi.Router
	httpServer *http.Server
	config     Config
}

//Run runs application server and other services
func (a *App) Run() {
	a.Logger.Info("Application started running")
	a.httpServer = &http.Server{
		Addr:    a.config.HTTPHostAddress,
		Handler: a.Router,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		a.Logger.Fatal("app: error while http listening: %s", zap.Error(err))
	}
	a.Logger.Fatal("http server closed with error: %s")
}

//Build is used to initialize application's dependencies
func (a *App) Build(initializer func() error) {
	err := initializer()
	if err != nil {
		a.Logger.Error("app: could not initialize application", zap.Error(err))
	}
}
