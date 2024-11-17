package database

import "app/models"

var Users []models.User = []models.User{
	{
		UserId:      "345f6guhio456789dfghjkl0",
		Firstname:   "Ali",
		Lastname:    "Umarov",
		PhoneNumber: "77",
		Password:    "0000",
		Subscribes: models.Subscribes{
			Convertor:  false,
			Calculator: false,
			Todo:       false,
		},
	},
	{
		UserId:      "345f6guhio456789dfghjkl1",
		Firstname:   "Vali",
		Lastname:    "Aliyev",
		PhoneNumber: "777",
		Password:    "000",
		Subscribes: models.Subscribes{
			Convertor:  false,
			Calculator: false,
			Todo:       false,
		},
	},
	{
		UserId:      "345f6guhio456789dfghjkl2",
		Firstname:   "Umar",
		Lastname:    "Valiyev",
		PhoneNumber: "7777",
		Password:    "00",
		Subscribes: models.Subscribes{
			Convertor:  false,
			Calculator: false,
			Todo:       true,
		},
	},
}
