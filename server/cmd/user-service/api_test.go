package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/stretchr/testify/require"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	mockbookly "github.com/pw-software-engineering/b-team/server/pkg/mocks/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const clientPath = "/api/v1/client/"

// this is simpler way of doing these tests, but it does not include routing
func Test_api_handleGetUser(t *testing.T) {
	mockErr := errors.New("mock err")
	exampleSession := bookly.Session{UserID: 1, HotelID: 0}
	type fields struct {
		logger      *zap.Logger
		userService *mockbookly.MockUserService
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, handler http.HandlerFunc)
	}{
		{
			name: "Wrong path should result in status not found",
			check: func(t *testing.T, handler http.HandlerFunc) {
				assert.HTTPStatusCode(t, handler, http.MethodGet, "/notImplemented", nil, http.StatusNotFound)
			},
		},
		{
			name: "if user service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(bookly.User{}, mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodGet, clientPath, nil, testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to get user")
			},
		},
		{
			name: "correct respond",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(bookly.User{}, nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodGet, clientPath, nil, testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:      zap.NewNop(),
				userService: mockbookly.NewMockUserService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.userService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}

func Test_api_handlePatchUser(t *testing.T) {
	mockErr := errors.New("mock err")
	exampleRequest := PatchUserRequest{}
	exampleSession := bookly.Session{UserID: 1, HotelID: 0}
	type fields struct {
		logger      *zap.Logger
		userService *mockbookly.MockUserService
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, handler http.HandlerFunc)
	}{
		{
			name: "Wrong path should result in status not found",
			check: func(t *testing.T, handler http.HandlerFunc) {
				assert.HTTPStatusCode(t, handler, http.MethodPatch, "/notImplemented", nil, http.StatusNotFound)
			},
		},
		{
			name: "header doesnt contain user session",
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPatch, clientPath, exampleRequest)
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "if user service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().UpdateUserInformation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPatch, clientPath, exampleRequest, testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				// todo: Check return struct
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
				req, err := http.NewRequest(http.MethodPatch, clientPath, &body)
				require.NoError(t, err)
				auth.SetSessionHeader(req.Header, &bookly.Session{})

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to update user info")
			},
		},
		{
			name: "correct respond",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().UpdateUserInformation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPatch, clientPath, exampleRequest)
				auth.SetSessionHeader(req.Header, &bookly.Session{})
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:      zap.NewNop(),
				userService: mockbookly.NewMockUserService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.userService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}

func Test_api_handlePostUserVerify(t *testing.T) {
	mockErr := errors.New("mock err")
	exampleRequest := PostUserRequest{}
	type fields struct {
		logger      *zap.Logger
		userService *mockbookly.MockUserService
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, handler http.HandlerFunc)
	}{
		{
			name: "Wrong path should result in status not found",
			check: func(t *testing.T, handler http.HandlerFunc) {
				assert.HTTPStatusCode(t, handler, http.MethodPost, "/notImplemented", nil, http.StatusNotFound)
			},
		},
		{
			name: "if user service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().UserVerify(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(false, bookly.Token{}, mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, clientPath+"login", exampleRequest)

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
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
				req, err := http.NewRequest(http.MethodPost, clientPath+"login", &body)

				require.NoError(t, err)

				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to login")
			},
		},
		{
			name: "correct respond with login accept",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().UserVerify(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(true, bookly.Token{}, nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, clientPath+"login", exampleRequest)
				auth.SetSessionHeader(req.Header, &bookly.Session{})
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusOK, recorder.Code)
				// todo: check token
			},
		},
		{
			name: "correct respond without login accept respond",
			prepare: func(t *testing.T, f *fields) {
				f.userService.EXPECT().UserVerify(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(false, bookly.Token{}, nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, clientPath+"login", exampleRequest)
				auth.SetSessionHeader(req.Header, &bookly.Session{})
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
				// todo: check token
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				logger:      zap.NewNop(),
				userService: mockbookly.NewMockUserService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.userService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}
