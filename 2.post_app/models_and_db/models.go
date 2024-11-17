package modelsanddb

import "time"

type User struct {
	UserId   string
	Username string
	Gmail    string
	Password string
}

type Post struct {
	PostId    string
	Title     string
	CreatedAt time.Time
	UserId    string // -----> Har bir postda foydalanuvchi id si bolishi kerak,  one to many bog`lanish !
}

type Login struct {
	Username string
	Password string
}
