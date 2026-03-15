package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"ecommerce/modules/auth"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	userHandler    *user.Handler
	authHandler    *auth.Handler
	productHandler *product.Handler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	authHandler *auth.Handler,
	productHandler *product.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		authHandler:    authHandler,
		productHandler: productHandler,
	}
}

func (server *Server) StartServer() {
	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	// getRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.authHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

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
