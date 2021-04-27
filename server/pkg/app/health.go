package app

import (
	"context"
	"fmt"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type health struct {
	configs []*gosundheit.Config
	health  gosundheit.Health
}

func newHealth(opts ...healthOpt) (*health, func()) {
	h := &health{
		configs: make([]*gosundheit.Config, 0),
		health:  gosundheit.New(),
	}

	for _, opt := range opts {
		opt(h)
	}
	return h, h.health.DeregisterAll
}

func (h *health) AddHealthCheck(cfg *gosundheit.Config) {
	h.configs = append(h.configs, cfg)
}

func (h *health) start() error {
	for _, config := range h.configs {
		if config == nil {
			return fmt.Errorf("app health.start: added health config cannot be nil")
		}
		if config.Check == nil {
			return fmt.Errorf("app health.start: added health check cannot be nil")
		}
		if details, err := config.Check.Execute(); err != nil {
			return fmt.Errorf("app health.start: initial health check from %s did not pass: err=%w details=%v", config.Check.Name(), err, details)
		}
		if err := h.health.RegisterCheck(config); err != nil {
			return fmt.Errorf("app health.start: could not register healtcheck %s: %w", config.Check.Name(), err)
		}
	}
	return nil
}

func (h *health) mount(router chi.Router) {
	router.Mount("/health", healthhttp.HandleHealthJSON(h.health))
}

type healthOpt func(h *health)

func withCancelOnUnhealthy(logger *zap.Logger, retries int64, cancelFunc context.CancelFunc) gosundheit.Option {
	return gosundheit.WithHealthListeners(&cancelOnUnhealthyListener{
		cancel:  cancelFunc,
		retries: retries,
		logger:  logger,
	})
}

func withGosundheitOpts(opts ...gosundheit.Option) healthOpt {
	return func(h *health) {
		h.health = gosundheit.New(opts...)
	}
}

type cancelOnUnhealthyListener struct {
	cancel  context.CancelFunc
	retries int64
	logger  *zap.Logger
}

var _ gosundheit.HealthListener = &cancelOnUnhealthyListener{}

func (c *cancelOnUnhealthyListener) OnResultsUpdated(results map[string]gosundheit.Result) {
	for _, result := range results {
		if result.ContiguousFailures > c.retries {
			c.logger.Info("health check failed too many times - canceling",
				zap.Any("failed_result", result),
			)
			c.cancel()
		}
	}
}
