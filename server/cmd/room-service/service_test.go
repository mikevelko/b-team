package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	mockbookly "github.com/pw-software-engineering/b-team/server/pkg/mocks/pkg/bookly"

	"github.com/stretchr/testify/assert"
)

func TestRoomService_CreateRoom(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
		room         *bookly.Room
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if room is not valid, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.room = &bookly.Room{RoomNumber: ""}
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.CreateRoom(nil, *f.room, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if user storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomStorage.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), mockErr)
				f.room = &bookly.Room{RoomNumber: "12"}
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.CreateRoom(nil, *f.room, 0)
				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomService_DeleteRoom(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if user storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomStorage.EXPECT().DeleteRoom(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.DeleteRoom(nil, 0, 0)
				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomService_GetAllHotelRooms(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if user storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomStorage.EXPECT().GetAllHotelRooms(gomock.Any(), gomock.Any()).Return(nil, mockErr)
				// f.roomStorage.EXPECT().GetRoom(gomock.Any(), gomock.Any(), gomock.Any()).Return(bookly.RoomError, bookly.Room{}, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetAllHotelRooms(nil, 0, 0, 0, "")
				assert.Error(t, err)
			},
		},
		{
			name: "if user storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.roomStorage.EXPECT().GetAllHotelRooms(gomock.Any(), gomock.Any()).Return(nil, mockErr)
				// f.roomStorage.EXPECT().GetRoom(gomock.Any(), gomock.Any(), gomock.Any()).Return(bookly.RoomError, bookly.Room{}, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetAllHotelRooms(nil, 0, 1, 1, "")
				assert.Error(t, err)
			},
		},
		{
			name: "if user storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				// f.roomStorage.EXPECT().GetAllHotelRooms(gomock.Any(), gomock.Any()).Return(nil, mockErr)
				f.roomStorage.EXPECT().GetRoomByName(gomock.Any(), gomock.Any(), gomock.Any()).Return(bookly.Room{}, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetAllHotelRooms(nil, 0, 0, 0, "12")
				assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomService_GetRoomsRelatedWithOffer(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if offer storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetRoomsRelatedWithOffer(nil, 0, 0, 0, 0, "")
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().RoomsRelatedWithOffer(gomock.Any(), gomock.Any()).Return(nil, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetRoomsRelatedWithOffer(nil, 0, 0, 0, 0, "")
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().RoomsRelatedWithOffer(gomock.Any(), gomock.Any()).Return([]int64{1, 2, 3}, nil)
				f.roomStorage.EXPECT().GetRoom(gomock.Any(), gomock.Any()).Return(bookly.Room{}, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetRoomsRelatedWithOffer(nil, 0, 0, 0, 0, "")
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().RoomsRelatedWithOffer(gomock.Any(), gomock.Any()).Return([]int64{1, 2, 3}, nil)
				f.roomStorage.EXPECT().GetRoom(gomock.Any(), gomock.Any()).Return(bookly.Room{}, nil)
				f.roomStorage.EXPECT().OffersRelatedWithRoom(gomock.Any(), gomock.Any()).Return(nil, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetRoomsRelatedWithOffer(nil, 0, 0, 0, 0, "")
				assert.Error(t, err)
			},
		},
		{
			name: "all is fine",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().RoomsRelatedWithOffer(gomock.Any(), gomock.Any()).Return([]int64{1}, nil)
				f.roomStorage.EXPECT().GetRoom(gomock.Any(), gomock.Any()).Return(bookly.Room{}, nil)
				f.roomStorage.EXPECT().OffersRelatedWithRoom(gomock.Any(), gomock.Any()).Return([]int64{1}, nil)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				_, err := service.GetRoomsRelatedWithOffer(nil, 0, 0, 0, 0, "")
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomService_AddRoomToOffer(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if offer storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.AddRoomToOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.AddRoomToOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsExistLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.AddRoomToOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsExistLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				f.roomStorage.EXPECT().AddLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.AddRoomToOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "all is fine",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsExistLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				f.roomStorage.EXPECT().AddLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.AddRoomToOffer(nil, 0, 0, 0)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomService_DeleteRoomFromOffer(t *testing.T) {
	mockErr := errors.New("mock err")
	type fields struct {
		roomStorage  *mockbookly.MockRoomStorage
		offerStorage *mockbookly.MockOfferStorage
	}
	tests := []struct {
		name    string
		prepare func(t *testing.T, f *fields)
		check   func(t *testing.T, service *roomService, f *fields)
	}{
		{
			name: "if offer storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.DeleteRoomFromOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.DeleteRoomFromOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "if room storage returns error, error is expected",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().DeleteLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockErr)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.DeleteRoomFromOffer(nil, 0, 0, 0)
				assert.Error(t, err)
			},
		},
		{
			name: "all is fine",
			prepare: func(t *testing.T, f *fields) {
				f.offerStorage.EXPECT().IsOfferOwnedByHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().IsRoomBelongToHotel(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
				f.roomStorage.EXPECT().DeleteLinkWithRoomAndOffer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			check: func(t *testing.T, service *roomService, f *fields) {
				err := service.DeleteRoomFromOffer(nil, 0, 0, 0)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := &fields{
				roomStorage:  mockbookly.NewMockRoomStorage(ctrl),
				offerStorage: mockbookly.NewMockOfferStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(t, f)
			}

			service := newRoomService(f.roomStorage, f.offerStorage)
			tt.check(t, service, f)
		})
	}
}

func TestRoomValid(t *testing.T) {
	tests := []struct {
		name  string
		args  bookly.Room
		check func(t *testing.T, err error)
	}{
		{
			name: "RoomNumber with wrong length",
			args: bookly.Room{RoomNumber: ""},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Valid RoomNumber",
			args: bookly.Room{RoomNumber: "12C"},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRoom(tt.args)
			tt.check(t, err)
		})
	}
}
