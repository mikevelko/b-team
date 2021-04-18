package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/require"

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

func Test_api_handlePatchManagerHotelDetails(t *testing.T) {
	exampleRequest := HotelPatchRequest{}
	exampleSession := bookly.Session{HotelID: 1}
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
		{
			name: "no hotel id in session returns error",
			prepare: func(t *testing.T, f *fields) {
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPatch, "/api/v1/hotel/hotelInfo", exampleRequest)
				auth.SetSessionHeader(req.Header, &bookly.Session{})

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "User is not a manager of any hotel")
			},
		},
		{
			name:    "when request doesn't have content type we expect error",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				err := json.NewEncoder(&body).Encode(exampleRequest)
				require.NoError(t, err)
				req, err := http.NewRequest(http.MethodPatch, "/api/v1/hotel/hotelInfo", &body)
				require.NoError(t, err)
				auth.SetSessionHeader(req.Header, &exampleSession)

				handler.ServeHTTP(recorder, req)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to update hotel details")
			},
		},
		{
			name:    "when server can't decode request we expect error",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				body.WriteString("{definitelyNotFakeJSONParameter:2}")

				req, err := http.NewRequest(http.MethodPatch, "/api/v1/hotel/hotelInfo", &body)
				require.NoError(t, err)

				req.Header.Set("Content-Type", "application/json")
				auth.SetSessionHeader(req.Header, &exampleSession)

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to update hotel details")
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
