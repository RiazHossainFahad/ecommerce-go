package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func getRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /products",
		manager.With( // will execute top to bottom
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.StoreProduct),
			middleware.Authenticate,
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.Authenticate,
		),
	)
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.Authenticate,
		),
	)

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.StoreUser),
		),
	)

	mux.Handle(
		"POST /login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}
