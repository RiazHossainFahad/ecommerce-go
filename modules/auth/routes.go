package auth

import (
	"ecommerce/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)
}
