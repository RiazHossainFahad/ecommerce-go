package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"ecommerce/middleware"
	"ecommerce/modules/auth"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
	"ecommerce/repo"
)

func main() {
	cnf := config.GetConfig()

	// Inject repository
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewareHanlder := middleware.NewMiddlewareHandler(
		cnf,
	)

	server := cmd.NewServer(
		cnf,
		user.NewHandler(userRepo),
		auth.NewHandler(
			cnf,
			userRepo,
		),
		product.NewHandler(productRepo, middlewareHanlder),
	)

	server.StartServer()
}
