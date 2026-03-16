package product

import (
	"ecommerce/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With( // will execute top to bottom
			http.HandlerFunc(h.GetProducts),
		))

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.StoreProduct),
			h.midlewareHandler.Authenticate,
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.midlewareHandler.Authenticate,
		),
	)
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.midlewareHandler.Authenticate,
		),
	)
}
