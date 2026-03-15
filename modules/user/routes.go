package user

import (
	"ecommerce/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.StoreUser),
		),
	)
}
