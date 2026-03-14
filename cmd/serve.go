package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	getRoutes(mux, manager)

	fmt.Printf("Server running on port %s\n", addr)

	// will execute top to bottom
	manager.Use(middleware.AddLog)
	manager.Use(middleware.HandleCors)
	manager.Use(middleware.Preflight)

	err := http.ListenAndServe(addr, manager.WrapMux(
		mux,
	)) // which port

	if err != nil {
		fmt.Println("Error while starting the server.\nError", err)
	}
}
