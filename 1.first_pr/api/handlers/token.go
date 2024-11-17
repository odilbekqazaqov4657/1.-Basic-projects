package handlers

import "app/models"

type Token struct {
	Username   string
	UserId     string
	Subscribes models.Subscribes
}

var UserToken *Token
