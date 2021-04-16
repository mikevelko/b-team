package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	mockbookly "github.com/pw-software-engineering/b-team/server/pkg/mocks/pkg/bookly"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_api_handleGetHotelPreviews(t *testing.T) {
	type fields struct {
		logger       *zap.Logger
		hotelService *mockbookly.MockHotelService
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, handler http.HandlerFunc)
	}{
		{
			name: "wrong path should result in status not found",
			check: func(t *testing.T, handler http.HandlerFunc) {
				assert.HTTPStatusCode(t, handler, http.MethodPatch, "/notImplemented", nil, http.StatusNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:       zap.NewNop(),
				hotelService: mockbookly.NewMockHotelService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.hotelService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}
