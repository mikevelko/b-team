package app

import (
	"context"
	"net/http"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"

	"github.com/go-chi/chi/middleware"

	"go.uber.org/zap/zapcore"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Config is used to configure application
type Config struct {
	HTTPHostAddress         string        `default:"0.0.0.0:8080" split_words:"true"`
	LogErrorsWithStacktrace bool          `default:"false" split_words:"true"`
	HealthInterval          time.Duration `default:"15s" split_words:"true"`
	// FailOnCheckRetries says, how many consecutive health check fails causes application to shutdown.
	//	negative number indicate, that application should never fail because of failed healthcheck
	FailOnCheckRetries int64 `default:"0" split_words:"true"`

	ShutdownTimeout time.Duration `default:"5s" split_words:"true"`
}

// App is responsible for initializing and running application
type App struct {
	Logger       *zap.Logger
	Router       chi.Router
	httpServer   *http.Server
	config       Config
	globalCtx    context.Context
	globalCancel context.CancelFunc
	*cleanupHandler
	*health
}

// NewApp initializes application
func NewApp(logger *zap.Logger, cfg Config) *App {
	if !cfg.LogErrorsWithStacktrace {
		logger.WithOptions(zap.AddStacktrace(zapcore.PanicLevel))
	}

	router := chi.NewRouter().
		With(
			middleware.DefaultLogger,
			middleware.Recoverer,
		)
	app := &App{
		Logger:         logger,
		Router:         router,
		config:         cfg,
		cleanupHandler: newCleanupHandler(logger),
	}
	app.globalCtx, app.globalCancel = context.WithCancel(context.Background())

	gosundheitOpts := []gosundheit.Option{}
	if cfg.FailOnCheckRetries >= 0 {
		gosundheitOpts = append(
			gosundheitOpts,
			withCancelOnUnhealthy(logger, cfg.FailOnCheckRetries, app.globalCancel),
		)
	}

	health, cleanup := newHealth(withGosundheitOpts(gosundheitOpts...))
	app.health = health
	app.AddCleanup(cleanup)
	return app
}

// Run runs application server and other services
func (a *App) Run() {
	a.Logger.Info("Application is starting running",
		zap.String("address", a.config.HTTPHostAddress))

	err := a.health.start()
	if err != nil {
		a.Logger.Fatal("app: error while starting health check", zap.Error(err))
	}
	a.httpServer = &http.Server{
		Addr:    a.config.HTTPHostAddress,
		Handler: a.Router,
	}

	// start http server
	go func() {
		err := a.httpServer.ListenAndServe() // always returns non-nil error
		if err == http.ErrServerClosed {     // purposeful shutdown
			a.Logger.Info("app: http server stopped")
			return
		}
		a.Logger.Fatal("app: error while http listening", zap.Error(err))
	}()
	a.listenForStop()
}

// listenForStop makes sure, that on graceful shutdown requests are handled before exiting
func (a *App) listenForStop() {
	<-a.globalCtx.Done()
	a.Logger.Info("app: attempting graceful shutdown")

	// graceful http shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	defer cancel()
	err := a.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		a.Logger.Info("app: error on graceful shutdown", zap.Error(err))
	}
	a.cleanup()
}

// Build is used to initialize application's dependencies
func (a *App) Build(initializer func() error) {
	err := initializer()
	if err != nil {
		a.Logger.Fatal("app: could not build application's dependencies", zap.Error(err))
		return
	}
	a.health.mount(a.Router)
}
