package models

import "time"

type Subscribes struct {
	Convertor  bool
	Calculator bool
	Todo       bool
}

type User struct {
	UserId      string
	Firstname   string
	Lastname    string
	PhoneNumber string
	Password    string
	Age         uint
	CreatedAt   time.Time
	Subscribes  Subscribes
}
