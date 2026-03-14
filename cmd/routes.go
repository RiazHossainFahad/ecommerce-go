package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func GetRoutes(mux *http.ServeMux) {
	// WITH GLOBAL ROUTER

	manager := middleware.NewManager()
	manager.Use(middleware.AddLog)

	// handler := http.HandlerFunc(handlers.GetProducts)
	// logMiddleware := middleware.AddLog(handler)

	// mux.Handle("GET /products", logMiddleware)

	// mux.Handle("GET /products", middleware.AddLog(http.HandlerFunc(handlers.GetProducts)))
	// mux.Handle("GET /products", middleware.NewManager().With(middleware.AddLog)(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.StoreProduct),
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
}
