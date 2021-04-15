package main

import (
	"encoding/json"
	"net/http"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"

	"github.com/go-chi/chi"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"go.uber.org/zap"
)

type api struct {
	logger      *zap.Logger
	userService bookly.UserService
}

func newAPI(logger *zap.Logger, service bookly.UserService) *api {
	return &api{
		logger:      logger,
		userService: service,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/client", func(r chi.Router) {
		r.Route("/login", func(r chi.Router) {
			r.Post("/", a.handlePostUserVerify)
		})
		r.With(auth.SessionMiddleware()).Route("/", func(r chi.Router) {
			r.Get("/", a.handleGetUser)
			r.Patch("/", a.handlePatchUser)
		})
	})
}

func (a *api) handleGetUser(w http.ResponseWriter, r *http.Request) {
	sesion := auth.SessionFromContext(r.Context())

	user, err := a.userService.GetUser(r.Context(), sesion.UserID)
	if err != nil {
		httputils.RespondWithError(w, "Unable to get user (Service error)")
		return
	}

	userResponse := GetUserResponseFromUser(user)

	httputils.WriteJSONResponse(a.logger, w, userResponse)
	return
}

func (a *api) handlePatchUser(w http.ResponseWriter, r *http.Request) {
	sesion := auth.SessionFromContext(r.Context())

	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to update user info (Wrong header)") {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decodedRequest PatchUserRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httputils.RespondWithError(w, "Unable to update user info (Decode error)")
		return
	}

	err = a.userService.UpdateUserInformation(r.Context(), decodedRequest.Username, decodedRequest.Email, sesion.UserID)
	if err != nil {
		// todo: Add better respond with error: Wrong email only or wrong username only
		pathUserRespond := PathUserErrorResponse{
			EmailError:    ErrorResponse{Message: err.Error()},
			UserNameError: ErrorResponse{Message: err.Error()},
		}
		// todo: StatusBadRequest could be change in spec !!!!!!!!
		errW := httputils.JSONRespondError(w, pathUserRespond, http.StatusBadRequest)
		if errW != nil {
			httputils.RespondWithError(w, "Unable to update user info (JSONErrorRespond error)")
			return
		}
		return
	}
	httputils.RespondWithCode(w, http.StatusOK)
	return
}

func (a *api) handlePostUserVerify(w http.ResponseWriter, r *http.Request) {
	if !httputils.IsHeaderTypeValid(w, r, "application/json", "Unable to login") {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decodedRequest PostUserRequest
	err := decoder.Decode(&decodedRequest)
	if err != nil {
		httputils.RespondWithError(w, "Unable to login")
		return
	}

	check, token, err := a.userService.UserVerify(r.Context(), decodedRequest.Login, decodedRequest.Password)
	if err != nil {
		httputils.RespondWithError(w, "Unable to login")
		return
	}

	if check {
		httputils.WriteJSONResponse(a.logger, w, token)
		return
	}
	// todo: Add better respond with user exists or no
	respond := PostUserErrorResponse{Desc: "This user might not exists or you write wrong email and password"}
	err = httputils.JSONRespondError(w, respond, http.StatusUnauthorized)
	if err != nil {
		httputils.RespondWithError(w, "Unable to login")
		return
	}
	return
}
