package api

import "github.com/go-chi/chi/v5"

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			//example => /api/v1/health
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignUpUser)
				r.Get("/login", api.handleLoginUser)
				r.Get("/logout", api.handleLogoutUser)
			})
		})
	})
}
