package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/middleware"
	"ecommerce/modules/auth"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
	"ecommerce/repo"
	"log"
)

func main() {
	cnf := config.GetConfig()

	dbClient, err := db.NewDBConnection(cnf)
	if err != nil {
		log.Println(err)
	}

	// Inject repository
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbClient)

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
