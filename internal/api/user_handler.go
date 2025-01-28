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
		if errors.Is(err, services.ErrDuplicatedEmailOrPassoword) {
			_ = jsonutils.EncodeJson(w, r, http.StatusConflict, map[string]string{"error": "email or username already exists"})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]string{"id": id.String()})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
