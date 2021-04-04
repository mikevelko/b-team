package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	mockbookly "github.com/pw-software-engineering/b-team/server/pkg/mocks/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// this is simpler way of doing these tests, but it does not include routing
func Test_api_handlePostOfferSimple(t *testing.T) {
	mockErr := errors.New("mock err")
	exampleRequest := CreateOfferRequest{
		Isactive:            false,
		Offertitle:          "dfsdfs",
		Maxguests:           2,
		Description:         "sdfsd",
		Offerpreviewpicture: "dsfsd",
	}
	type fields struct {
		logger       *zap.Logger
		offerService *mockbookly.MockOfferService
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
			name: "if offer service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerService.EXPECT().
					HandleCreateOffer(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(0), mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, "/api/v1/hotel/offers", exampleRequest)
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to add offer")
			},
		},
		{
			name:    "when request doesn't have content type we expect status 400",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				err := json.NewEncoder(&body).Encode(exampleRequest)
				require.NoError(t, err)
				req, err := http.NewRequest(http.MethodPost, "/api/v1/hotel/offers", &body)

				require.NoError(t, err)

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to add offer")
			},
		},
		{
			name:    "when server can't decode request",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				body.WriteString("xcsdnvursvndjv")

				req, err := http.NewRequest(http.MethodPost, "/api/v1/hotel/offers", &body)
				require.NoError(t, err)

				req.Header.Set("Content-Type", "application/json")

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to add offer")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:       zap.NewNop(),
				offerService: mockbookly.NewMockOfferService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.offerService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}

func Test_api_handleGetOffers(t *testing.T) {
	mockErr := errors.New("mock err")
	exampleRequest := GetOffersRequest{}
	type fields struct {
		logger       *zap.Logger
		offerService *mockbookly.MockOfferService
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
			name: "if offer service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerService.EXPECT().
					HandleGetOffers(gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]bookly.Offer{}, mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodGet, "/api/v1/hotel/offers", exampleRequest)
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to get offers")
			},
		},
		{
			name:    "when request doesn't have content type we expect status 400",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				err := json.NewEncoder(&body).Encode(exampleRequest)
				require.NoError(t, err)
				req, err := http.NewRequest(http.MethodGet, "/api/v1/hotel/offers", &body)

				require.NoError(t, err)

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to get offers")
			},
		},
		{
			name:    "when server can't decode request",
			prepare: nil,
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()

				var body bytes.Buffer
				body.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit, " +
					"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." +
					" Ut enim ad minim veniam, quis nostrud exercitation ullamco" +
					" laboris nisi ut aliquip ex ea commodo consequat. " +
					"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu " +
					"fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in " +
					"culpa qui officia deserunt mollit anim id est laborum.")

				req, err := http.NewRequest(http.MethodGet, "/api/v1/hotel/offers", &body)
				require.NoError(t, err)

				req.Header.Set("Content-Type", "application/json")

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to get offers")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:       zap.NewNop(),
				offerService: mockbookly.NewMockOfferService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.offerService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}