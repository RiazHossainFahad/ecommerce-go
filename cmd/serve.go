package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	getRoutes(mux, manager)

	fmt.Println("Server running on port :3000")

	// will execute top to bottom
	manager.Use(middleware.AddLog)
	manager.Use(middleware.HandleCors)
	manager.Use(middleware.Preflight)

	err := http.ListenAndServe(":3000", manager.WrapMux(
		mux,
	)) // which port

	if err != nil {
		fmt.Println("Error while starting the server.\nError", err)
	}
}
