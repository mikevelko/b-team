package app

import (
    "go.uber.org/zap"
	"github.com/go-chi/chi"
    "net/http"
)

type Config struct{
    HttpHostAddress string `env:"default=0.0.0.0:8080"`
}



func NewApp(logger *zap.Logger, cfg Config) *App{
    return &App{
        Logger: logger,
        Router: chi.NewRouter(),
        config: cfg,
    }
}


// App is responsible for initializing and running application
type App struct{
    Logger     *zap.Logger
    Router     chi.Router
    httpServer *http.Server
    config Config
}

func (a *App) Run() {
    a.Logger.Info("Application started running")
    a.httpServer = &http.Server{
        Addr:    a.config.HttpHostAddress,
        Handler: a.Router,
    }

    if err := a.httpServer.ListenAndServe(); err != nil {
        a.Logger.Fatal( "app: error while http listening: %s", zap.Error(err))
    }
    a.Logger.Fatal("http server closed with error: %s")
}

func (a *App) Build(initializer func() error){
    err := initializer()
    if err != nil{
        a.Logger.Error("app: could not initialize application", zap.Error(err))
    }
}
