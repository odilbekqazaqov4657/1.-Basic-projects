package modelsanddb

import "time"

var (
	Users []User = []User{
		{
			UserId:   "bc3fb45a-24d9-4bc2-a0e4-745ca3bf79e7",
			Username: "Ali",
			Gmail:    "Ali001@gmail.com",
			Password: "Ali001",
		},

		{
			UserId:   "bc3fb45a-24d9-4bc2-a0e4-745ca3bf79e8",
			Username: "Vali",
			Gmail:    "Vali002@gmail.com",
			Password: "Vali002",
		},
		{
			UserId:   "bc3fb45a-24d9-4bc2-a0e4-745ca3bf79e9",
			Username: "Odilbek",
			Gmail:    "odilbekqazaqov4657@gmail.com",
			Password: "0777",
		},
	}

	Posts []Post = []Post{
		{
			PostId:    "5213a60c-c30a-4a44-ba09-93fbc0ccdeb1",
			Title:     "Picnic time !",
			CreatedAt: time.Time{},
			UserId:    "bc3fb45a-24d9-4bc2-a0e4-745ca3bf79e9",
		},
	}
)
