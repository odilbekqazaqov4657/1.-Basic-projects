package storage

import (
	modelsanddb "app/models_and_db"
	"errors"
)

type UserRepo struct {
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (u UserRepo) Login(login modelsanddb.Login) (user *modelsanddb.User, err error) {
	for _, u := range modelsanddb.Users {
		if u.Username == login.Username && u.Password == login.Password {

			user = &u

			return
		}

	}
	return nil, errors.New("you are not registered or no account with that username ")
}
