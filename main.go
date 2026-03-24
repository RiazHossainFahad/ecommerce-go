package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/middleware"
	authHandler "ecommerce/modules/auth"
	prdctHandler "ecommerce/modules/product"
	usrHandler "ecommerce/modules/user"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/user"
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

	// Domains
	prodctDomain := product.NewService(productRepo)
	userDomain := user.NewService(userRepo)

	middlewareHanlder := middleware.NewMiddlewareHandler(
		cnf,
	)

	server := cmd.NewServer(
		cnf,
		usrHandler.NewHandler(userDomain),
		authHandler.NewHandler(
			cnf,
			userDomain,
		),
		prdctHandler.NewHandler(prodctDomain, middlewareHanlder),
	)

	server.StartServer()
}
