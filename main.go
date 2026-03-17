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
	"fmt"
	"log"
	"os"
)

func main() {
	cnf := config.GetConfig()

	// fmt.Printf("%+v", cnf)

	dbClient, err := db.NewDBConnection(cnf)
	if err != nil {
		log.Println(err)
	}

	err = db.MigrateDB(dbClient, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Inject repository
	productRepo := repo.NewProductRepo(dbClient)
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
