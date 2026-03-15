package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"ecommerce/modules/auth"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
)

func main() {
	cnf := config.GetConfig()
	server := cmd.NewServer(
		cnf,
		user.NewHandler(),
		auth.NewHandler(),
		product.NewHandler(cnf),
	)

	server.StartServer()
}
