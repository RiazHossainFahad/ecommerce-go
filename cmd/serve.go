package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // router

	// PREVIOUSLY AS BELOW
	// mux.HandleFunc("/products", handleCorsMiddleware(getProducts))
	// mux.HandleFunc("/store", handleCorsMiddleware(storeProduct))

	// WITH NORMAL HandlerFunc
	// mux.Handle("GET /products", http.HandlerFunc(getProducts))

	// WITH MIDDLEWARE
	// mux.Handle("GET /products", handleCorsMiddleware(http.HandlerFunc(getProducts)))
	// mux.Handle("POST /store", handleCorsMiddleware(http.HandlerFunc(storeProduct)))

	GetRoutes(mux)

	fmt.Println("Server running on port :3000")

	// err := http.ListenAndServe(":3000", mux) // which port
	err := http.ListenAndServe(":3000", middleware.HandleCorsMiddleware(mux)) // which port

	if err != nil {
		fmt.Println("Error while starting the server.\nError", err)
	}
}
