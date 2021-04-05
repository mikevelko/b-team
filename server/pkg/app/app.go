package app

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Config is used to configure application
type Config struct {
	HTTPHostAddress         string `default:"0.0.0.0:8080" split_words:"true"`
	LogErrorsWithStacktrace bool   `default:"false" split_words:"true"`
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

	return &App{
		Logger: logger,
		Router: router,
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

// Run runs application server and other services
func (a *App) Run() {
	a.Logger.Info("Application started running",
		zap.String("address", a.config.HTTPHostAddress))
	a.httpServer = &http.Server{
		Addr:    a.config.HTTPHostAddress,
		Handler: a.Router,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		a.Logger.Fatal("app: error while http listening", zap.Error(err))
	}
}

// Build is used to initialize application's dependencies
func (a *App) Build(initializer func() error) {
	err := initializer()
	if err != nil {
		a.Logger.Error("app: could not initialize application", zap.Error(err))
	}
}

// LoadConfig loads configuration from envs and prints usage if requirements are not fulfilled
func LoadConfig(logger *zap.Logger, cfg interface{}) {
	err := envconfig.Process("", cfg)
	if err != nil {
		envconfig.Usage("", cfg)
		logger.Fatal("app: could not process config", zap.Error(err))
	}
}
