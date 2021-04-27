package app

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestApp_HandleShutdown checks, if the application does correct graceful shutdown, when it fails
// 	1. start application with healthcheck passing
// 	2. do http request A which takes some time
// 	3. during that time do graceful unhealthy
// 	4. try doing request B
// 	5. A should finish successfully, B should fail
func TestApp_HandleShutdown(t *testing.T) {
	testTimeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	firstReqDoneErr := errors.New("first request was made, so the app is unhealthy")
	t.Cleanup(cancel)
	const port = "34542"

	firstReq, unhealthy := make(chan struct{}, 0), make(chan struct{}, 0)
	setUnhealthy := sync.Once{}

	cfg := Config{
		HTTPHostAddress:         "0.0.0.0:" + port,
		LogErrorsWithStacktrace: false,
		HealthInterval:          100 * time.Millisecond,
		FailOnCheckRetries:      0,
		ShutdownTimeout:         1 * time.Second,
	}

	logger := zap.NewNop()
	app := NewApp(logger, cfg)

	checkCfg := &gosundheit.Config{
		Check: testCheckFunc(func() error {
			select {
			case <-firstReq:
				setUnhealthy.Do(func() {
					close(unhealthy)
				})
				return firstReqDoneErr
			case <-testTimeoutCtx.Done():
				t.Fatal("test reached timeout!")
			default:
				return nil
			}
			return nil
		}),
		ExecutionPeriod: 100 * time.Millisecond,
	}
	app.Build(func() error {
		app.AddHealthCheck(checkCfg)
		isFirst := true // it's safe, because we do requests one after another
		app.Router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			if isFirst {
				isFirst = false
				close(firstReq)
				writer.WriteHeader(http.StatusOK)
				return
			}
			t.Fatal("second request is being handled, but it should never be possible")
		})
		return nil
	})

	go app.Run()

	r1, err := http.NewRequestWithContext(testTimeoutCtx, http.MethodGet, "http://localhost:"+port, http.NoBody)
	require.NoError(t, err)
	resp1, err := http.DefaultClient.Do(r1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)

	select {
	case <-unhealthy:
		select {
		case <-app.globalCtx.Done():
			r2, err := http.NewRequestWithContext(testTimeoutCtx, http.MethodGet, "localhost:"+port, http.NoBody)
			require.NoError(t, err)
			_, err = http.DefaultClient.Do(r2)
			assert.Error(t, err)
		case <-testTimeoutCtx.Done():
			t.Fatal("test reached timeout!")
		}
	case <-testTimeoutCtx.Done():
		t.Fatal("test reached timeout!")
	}
}
