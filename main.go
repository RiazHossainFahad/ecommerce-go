package main

import (
	"ecommerce/util"
	"log"
)

func main() {
	jwt, err := util.CreateJWT("my-secret", util.Payload{
		Sub:         1,
		FirstName:   "Fahad",
		LastName:    "Riaz",
		Email:       "fahad@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(jwt)
	// cmd.Serve()
}
