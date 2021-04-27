package app

import (
	"errors"
	"testing"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCheckFunc func() error

var _ checks.Check = testCheckFunc(func() error { return nil })

const (
	checkFuncDetail = "detail"
	checkFuncName   = "name"
)

func (t testCheckFunc) Name() string {
	return checkFuncName
}

func (t testCheckFunc) Execute() (interface{}, error) {
	return checkFuncDetail, t()
}

func TestHealth(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		checks []*gosundheit.Config
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, h *health)
	}{
		{
			name: "when there are no checks, health should not fail",
			prepare: func(t *testing.T, f *fields) {
				f.checks = []*gosundheit.Config{}
			},
			check: func(t *testing.T, h *health) {
				assert.NoError(t, h.start())
			},
		},
		{
			name: "when the check is nil, start should fail",
			prepare: func(t *testing.T, f *fields) {
				f.checks = []*gosundheit.Config{
					nil,
				}
			},
			check: func(t *testing.T, h *health) {
				assert.Error(t, h.start())
			},
		},
		{
			name: "when initial check returns error, start should fail and have init check error wrapped and contain detail",
			prepare: func(t *testing.T, f *fields) {
				f.checks = []*gosundheit.Config{
					{
						Check: testCheckFunc(func() error {
							return mockErr
						}),
					},
				}
			},
			check: func(t *testing.T, h *health) {
				err := h.start()
				require.Error(t, err)
				assert.ErrorIs(t, h.start(), mockErr)
				assert.Contains(t, err.Error(), checkFuncDetail)
				assert.Contains(t, err.Error(), checkFuncName)
			},
		},
		{
			name: "when second check fails and check is set to be run initially, start should not fail, buf status should be set to unhealthy soon after registering",
			prepare: func(t *testing.T, f *fields) {
				i := 1
				f.checks = []*gosundheit.Config{
					{
						Check: testCheckFunc(func() error {
							if i < 2 {
								i++
								return nil
							}
							return mockErr
						}),
						ExecutionPeriod:  10,
						InitialDelay:     0,
						InitiallyPassing: true,
					},
				}
			},
			check: func(t *testing.T, h *health) {
				require.NoError(t, h.start())
				time.Sleep(100 * time.Millisecond) // wait for that initial check from package. It's pretty nasty, but it should work under reasonable circumstances
				assert.False(t, h.health.IsHealthy())
			},
		},
		{
			name: "Happy path, no error at start and no error at init",
			prepare: func(t *testing.T, f *fields) {
				f.checks = []*gosundheit.Config{
					{
						Check: testCheckFunc(func() error {
							return nil
						}),
						ExecutionPeriod:  10,
						InitialDelay:     0,
						InitiallyPassing: false,
					},
				}
			},
			check: func(t *testing.T, h *health) {
				require.NoError(t, h.start())
				time.Sleep(100 * time.Millisecond) // wait for that initial check from package. It's pretty nasty, but it should work under reasonable circumstances
				assert.True(t, h.health.IsHealthy())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, cleanup := newHealth()
			t.Cleanup(cleanup)
			f := fields{}
			tt.prepare(t, &f)
			for _, check := range f.checks {
				h.AddHealthCheck(check)
			}
			tt.check(t, h)
		})
	}
}
