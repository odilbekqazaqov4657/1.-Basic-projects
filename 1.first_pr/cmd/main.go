package main

import (
	"app/api"
	"app/services"
	"app/storage/db"
)

func main() {

	userRepo := db.NewUserRepo()

	s := services.NewService()

	api.Api(userRepo, s)
}
