package http

import "github.com/go-chi/chi"

func (h *handler) AssignRoute(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
