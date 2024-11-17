package cache

import (
	"app/database"
	"app/models"
	"errors"
	"time"
)

var Caches = make(map[int]models.User)

type Cache struct {
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Set(otp int, user models.User) error {

	for _, e := range database.Users {
		if e.PhoneNumber == user.PhoneNumber {
			return errors.New("This phone number is already registered !")
		}
	}

	_, ok := Caches[otp]
	if ok {
		return errors.New("This phone number is already registered !")
	}

	Caches[otp] = user

	return nil
}

func (c Cache) Get(otp int, phoneNumber string) (*models.User, error) {

	_, ok := Caches[otp]
	if !ok {
		return nil, errors.New("Wrong otp !")
	}

	user := Caches[otp]

	user.CreatedAt = time.Now()
	if user.PhoneNumber != phoneNumber {
		return nil, errors.New("Wrong otp !")
	}

	delete(Caches, otp)

	return &user, nil
}
