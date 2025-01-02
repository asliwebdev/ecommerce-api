package main

import (
	"log"

	"ecommerce/handler"
	"ecommerce/postgres"
	"ecommerce/repository"
	"ecommerce/service"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo)

	h := handler.NewHandler(userService, productService)

	r := handler.Run(h)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
