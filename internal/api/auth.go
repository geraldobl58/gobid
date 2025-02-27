package api

import (
	"net/http"

	"github.com/geraldobl58/gobid/internal/jsonutils"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AutenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "must be logged in",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
