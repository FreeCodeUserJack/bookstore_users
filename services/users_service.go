package services

import (
	"github.com/FreeCodeUserJack/bookstore_users/domain/users"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
)


func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("userId cannot be <= 0", "bad_request")
	}

	return users.GetUserById(userId)
}