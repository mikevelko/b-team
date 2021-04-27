package main

import (
	//"encoding/json"
	//"net/http"

	//"github.com/pw-software-engineering/b-team/server/pkg/auth"

	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httputils"

	//"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"go.uber.org/zap"
)

type api struct {
	logger      *zap.Logger
	roomService bookly.RoomService
}

func newAPI(logger *zap.Logger, service bookly.RoomService) *api {
	return &api{
		logger:      logger,
		roomService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/hotel", func(r chi.Router) {
		r.With(auth.SessionMiddleware()).Route("/rooms", func(r chi.Router) {
			r.Get("/", a.handleGetRooms)
			r.Post("/", a.handlePostRoom)
			r.Delete("/{roomID}", a.handleDeleteRoom)
		})
	})
}

func (a *api) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to get room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httputils.RespondWithError(w, "User is not a manager of any hotel")
		return
	}

	pageNumberStr := r.URL.Query().Get("pageNumber")
	offersPerPageStr := r.URL.Query().Get("pageSize")
	roomNumber := r.URL.Query().Get("roomNumber")
	pageNumber, errPN := strconv.ParseInt(pageNumberStr, 10, 32)
	if errPN != nil {
		pageNumber = 1
	}
	pageSize, errRPP := strconv.ParseInt(offersPerPageStr, 10, 32)
	if errRPP != nil {
		pageSize = 10
	}

	a.logger.Info("Get Rooms", zap.Int64("pageNumber", pageNumber), zap.Int64("itemsPerPage", pageSize))

	roomsPreviews, err := a.roomService.GetAllHotelRooms(r.Context(), session.HotelID, int(pageNumber), int(pageSize), roomNumber)
	if err != nil {
		if errors.Is(err, bookly.ErrRoomNotFound) {
			httputils.RespondWithCode(w, http.StatusNotFound)
		} else {
			httputils.RespondWithError(w, "Unable to get room (Service Error)")
		}
		return
	}
	httputils.WriteJSONResponse(a.logger, w, roomsPreviews)
	return
}

func (a *api) handlePostRoom(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to post room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httputils.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to add room (Header Type Invalid)") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest string
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httputils.RespondWithError(w, "Unable to add room (")
		a.logger.Info("handlePostRoom: could not decode", zap.Error(err))
		return
	}
	room := bookly.Room{RoomNumber: decodedRequest}
	id, err := a.roomService.CreateRoom(r.Context(), room, session.HotelID)
	if err != nil {
		if errors.Is(err, bookly.ErrRoomAlreadyExists) {
			httputils.RespondWithCode(w, http.StatusConflict)
		} else {
			httputils.RespondWithError(w, "Unable to add room (Service Error)")
			a.logger.Info("handlePostOffer: could error create room", zap.Error(err))
		}
		return
	}
	httputils.WriteJSONResponse(a.logger, w, id)
	return
}

func (a *api) handleDeleteRoom(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to delete room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httputils.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	roomIDStr := chi.URLParam(r, "roomID")
	roomID, errConvert := strconv.ParseInt(roomIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete room, due bad parameter", zap.Error(errConvert))
		httputils.RespondWithError(w, "Unable to delete room (Conv Error)")
		return
	}

	err := a.roomService.DeleteRoom(r.Context(), roomID, session.HotelID)
	if err != nil {
		a.logger.Info("Unable to delete offer, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrRoomNotFound) {
			httputils.RespondWithCode(w, http.StatusNotFound)
		} else if errors.Is(err, bookly.ErrRoomNotBelongToHotel) {
			httputils.RespondWithCode(w, http.StatusUnauthorized)

			// todo: Add Error if room is occupy
		} else {
			httputils.RespondWithError(w, "Unable to delete room (Service Error)")
		}
		return
	}
	httputils.RespondWithCode(w, http.StatusOK)
	return
}
