package http

import "github.com/go-chi/chi"

func (h *handler) AssignRoute(r *chi.Mux) {
	r.Route("/process", func(r chi.Router) {
		r.Post("/payment", h.ProcessPayment)
	})
}
