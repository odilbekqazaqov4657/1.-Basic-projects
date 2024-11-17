package handlers

import (
	"app/services"
	"app/storage/db"
)

type Handler struct {
	userRepo db.UserRepo
	Service  services.Service
}

func NewHandler(userRepo db.UserRepo, service services.Service) Handler {
	h := Handler{
		userRepo: userRepo,
		Service:  service,
	}
	return h
}
