package api

import (
	"errors"
	"net/http"

	"github.com/geraldobl58/gobid/internal/jsonutils"
	"github.com/geraldobl58/gobid/internal/services"
	"github.com/geraldobl58/gobid/internal/usecase/user"
)

func (api *Api) handleSignUpUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(), data.UserName, data.Email, data.Password, data.Bio)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusConflict, map[string]string{"error": "email or username already exists"})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]string{"id": id.String()})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
	}

	id, err := api.UserService.AuthenticatedUser(r.Context(), data.Email, data.Password)

	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": "invalid credentials"})
			return
		}

		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	err = api.Sessions.RenewToken(r.Context())

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Put(r.Context(), "AutenticatedUserId", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "Logged in, user authenticated",
	})
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Remove(r.Context(), "AutenticatedUserId")

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "User logged out",
	})
}
