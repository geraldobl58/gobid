package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			//example => /api/v1/health
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignUpUser)
				r.Post("/login", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogoutUser)
			})
		})
	})
}
