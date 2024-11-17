package db

import (
	"app/database"
	"app/models"
	"errors"
)

type UserRepo struct {
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (u *UserRepo) CreateUser(user models.User) error {

	database.Users = append(database.Users, user)

	return nil
}

func (u *UserRepo) GetUsers() []models.User {

	return database.Users

}

func (u *UserRepo) Login(password, phoneNumber string) (*models.User, error) {
	for _, e := range database.Users {

		if e.Password == password && e.PhoneNumber == phoneNumber {
			return &e, nil
		}
	}
	return nil, errors.New("you do not have an accaunt: Create an accaunt ")
}
