package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	mockbookly "github.com/pw-software-engineering/b-team/server/pkg/mocks/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const roomPath = "/api/v1/hotel/rooms/"

// this is simpler way of doing these tests, but it does not include routing
func Test_api_handleGetRoom(t *testing.T) {
	mockErr := errors.New("mock err")
	// wrong_exampleSession := bookly.Session{UserID: 1, HotelID: 0}
	exampleSession := bookly.Session{UserID: 1, HotelID: 1}
	type fields struct {
		logger      *zap.Logger
		roomService *mockbookly.MockRoomService
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
			name: "if room service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().GetAllHotelRooms(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodGet, roomPath, nil, testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to get room")
			},
		},
		{
			name: "correct respond",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().GetAllHotelRooms(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]*bookly.Room{{
						ID:         1,
						RoomNumber: "12",
						HotelID:    1,
					}, {
						ID:         2,
						RoomNumber: "13",
						HotelID:    1,
					}}, nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodGet, roomPath, nil, testutils.WithSessionHeader(exampleSession))
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
				roomService: mockbookly.NewMockRoomService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.roomService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}

func Test_api_handlePostRoom(t *testing.T) {
	mockErr := errors.New("mock err")
	// wrong_exampleSession := bookly.Session{UserID: 1, HotelID: 0}
	exampleSession := bookly.Session{UserID: 1, HotelID: 1}
	type fields struct {
		logger      *zap.Logger
		roomService *mockbookly.MockRoomService
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
			name: "if room service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(0), mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, roomPath, "12", testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to add room")
			},
		},
		{
			name: "correct respond",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(1), nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodPost, roomPath, "12", testutils.WithSessionHeader(exampleSession))
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
				roomService: mockbookly.NewMockRoomService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.roomService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}

func Test_api_handleDeleteRoom(t *testing.T) {
	mockErr := errors.New("mock err")
	// wrong_exampleSession := bookly.Session{UserID: 1, HotelID: 0}
	exampleSession := bookly.Session{UserID: 1, HotelID: 1}
	type fields struct {
		logger      *zap.Logger
		roomService *mockbookly.MockRoomService
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, handler http.HandlerFunc)
	}{
		{
			name: "Wrong path should result in status not found",
			check: func(t *testing.T, handler http.HandlerFunc) {
				assert.HTTPStatusCode(t, handler, http.MethodDelete, "/notImplemented", nil, http.StatusNotFound)
			},
		},
		{
			name: "if room service returns error, internal server error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().DeleteRoom(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(mockErr)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodDelete, roomPath+"1", nil, testutils.WithSessionHeader(exampleSession))
				handler.ServeHTTP(recorder, req)
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				resp := testutils.ErrRespFromBody(t, recorder.Body)
				assert.Contains(t, resp.Error, "Unable to delete room")
			},
		},
		{
			name: "correct respond",
			prepare: func(t *testing.T, f *fields) {
				f.roomService.EXPECT().DeleteRoom(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			check: func(t *testing.T, handler http.HandlerFunc) {
				recorder := httptest.NewRecorder()
				req := testutils.JSONRequest(t, http.MethodDelete, roomPath+"1", nil, testutils.WithSessionHeader(exampleSession))
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
				roomService: mockbookly.NewMockRoomService(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			api := newAPI(f.logger, f.roomService)
			router := chi.NewRouter()
			api.mount(router)
			tt.check(t, router.ServeHTTP)
		})
	}
}
