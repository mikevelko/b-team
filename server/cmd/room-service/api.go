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
	"github.com/pw-software-engineering/b-team/server/pkg/httpapi"

	//"github.com/pw-software-engineering/b-team/server/pkg/httpapi"
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
		r.With(auth.SessionMiddleware()).Route("/offers/{offerID}/rooms", func(r chi.Router) {
			r.Get("/", a.handleGetRoomsRelatedWithOffer)
			r.Post("/", a.handlePostRoomToOffer)
			r.Delete("/{roomID}", a.handleDeleteRoomFromOffer)
		})
	})
}

func (a *api) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to get room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
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
			httpapi.RespondWithCode(w, http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to get room (Service Error)")
		}
		return
	}
	httpapi.WriteJSONResponse(a.logger, w, roomsPreviews)
	return
}

func (a *api) handlePostRoom(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to post room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to add room (Header Type Invalid)") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest RoomRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add room (")
		a.logger.Info("handlePostRoom: could not decode", zap.Error(err))
		return
	}
	room := bookly.Room{RoomNumber: decodedRequest.HotelRoomNumber}
	id, err := a.roomService.CreateRoom(r.Context(), room, session.HotelID)
	if err != nil {
		if errors.Is(err, bookly.ErrRoomAlreadyExists) {
			httpapi.RespondWithCode(w, http.StatusConflict)
		} else {
			httpapi.RespondWithError(w, "Unable to add room (Service Error)")
			a.logger.Info("handlePostOffer: could error create room", zap.Error(err))
		}
		return
	}
	var resp RoomRespond
	resp.RoomID = id
	httpapi.WriteJSONResponse(a.logger, w, resp)
	return
}

func (a *api) handleDeleteRoom(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to delete room, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	roomIDStr := chi.URLParam(r, "roomID")
	roomID, errConvert := strconv.ParseInt(roomIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete room, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithError(w, "Unable to delete room (Conv Error)")
		return
	}

	err := a.roomService.DeleteRoom(r.Context(), roomID, session.HotelID)
	if err != nil {
		a.logger.Info("Unable to delete room, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrRoomNotFound) {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		} else if errors.Is(err, bookly.ErrRoomNotOwnedByHotel) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)

			// todo: Add Error if room is occupy by reservation
		} else {
			httpapi.RespondWithError(w, "Unable to delete room (Service Error)")
		}
		return
	}
	httpapi.RespondWithCode(w, http.StatusOK)
	return
}

func (a *api) handleGetRoomsRelatedWithOffer(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to get rooms related with offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}

	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to get room related with offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithError(w, "Unable to get room related with offer from offer (Conv Error)")
		return
	}

	pageNumberStr := r.URL.Query().Get("pageNumber")
	offersPerPageStr := r.URL.Query().Get("pageSize")
	roomNumber := r.URL.Query().Get("roomNumber")
	pageNumber, err := strconv.ParseInt(pageNumberStr, 10, 32)
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.ParseInt(offersPerPageStr, 10, 32)
	if err != nil {
		pageSize = 10
	}

	a.logger.Info("Get Rooms From Offer", zap.Int64("pageNumber", pageNumber), zap.Int64("itemsPerPage", pageSize))

	roomsPreviews, err := a.roomService.GetRoomsRelatedWithOffer(r.Context(), offerID, session.HotelID, int(pageNumber), int(pageSize), roomNumber)
	if err != nil {
		if errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else if errors.Is(err, bookly.ErrRoomNotFound) {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to get rooms related with offer (Service Error)")
		}
		return
	}
	httpapi.WriteJSONResponse(a.logger, w, roomsPreviews)
	return
}

func (a *api) handlePostRoomToOffer(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to post room to offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}

	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete room from offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithError(w, "Unable to delete room from offer (Conv Error)")
		return
	}

	if !httpapi.IsHeaderTypeValid(w, r, "application/json", "Unable to add room to offer (Header Type Invalid)") {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var decodedRequest int
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httpapi.RespondWithError(w, "Unable to add room to offer")
		a.logger.Info("handlePostRoomToOffer: could not decode", zap.Error(err))
		return
	}

	err = a.roomService.AddRoomToOffer(r.Context(), offerID, int64(decodedRequest), session.HotelID)
	if err != nil {
		a.logger.Info("Unable to add room from offer, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrRoomNotOwnedByHotel) || errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else if errors.Is(err, bookly.ErrLinkOfferRoomNotFound) {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to add room from offer (Service Error)")
		}
		return
	}
	httpapi.RespondWithCode(w, http.StatusOK)
	return
}

func (a *api) handleDeleteRoomFromOffer(w http.ResponseWriter, r *http.Request) {
	session := auth.SessionFromContext(r.Context())
	if session.HotelID == 0 {
		a.logger.Info("Unable to delete room from offer, since logged person doesnt have assigned hotel", zap.Int64("UserID", session.UserID))
		httpapi.RespondWithError(w, "User is not a manager of any hotel")
		return
	}
	roomIDStr := chi.URLParam(r, "roomID")
	roomID, errConvert := strconv.ParseInt(roomIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete room from offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithError(w, "Unable to delete room from offer (Conv Error)")
		return
	}
	offerIDStr := chi.URLParam(r, "offerID")
	offerID, errConvert := strconv.ParseInt(offerIDStr, 10, 64)
	if errConvert != nil {
		a.logger.Info("Unable to delete room from offer, due bad parameter", zap.Error(errConvert))
		httpapi.RespondWithError(w, "Unable to delete room from offer (Conv Error)")
		return
	}

	err := a.roomService.DeleteRoomFromOffer(r.Context(), offerID, roomID, session.HotelID)
	if err != nil {
		a.logger.Info("Unable to delete room from offer, due to internal server error", zap.Error(err))
		if errors.Is(err, bookly.ErrRoomNotOwnedByHotel) || errors.Is(err, bookly.ErrOfferNotOwned) {
			httpapi.RespondWithCode(w, http.StatusUnauthorized)
		} else if errors.Is(err, bookly.ErrLinkOfferRoomNotFound) {
			httpapi.RespondWithCode(w, http.StatusNotFound)
		} else {
			httpapi.RespondWithError(w, "Unable to delete room from offer (Service Error)")
		}
		return
	}
	httpapi.RespondWithCode(w, http.StatusOK)
	return
}
