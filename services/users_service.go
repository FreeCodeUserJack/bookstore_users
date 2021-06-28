package services

import (
	"strings"

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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	savedUser, err := GetUserById(user.Id)
	if err != nil {
		return nil, err
	}

	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Lastname = strings.TrimSpace(user.Lastname)
	user.Email = strings.TrimSpace(user.Email)
	user.Status = strings.TrimSpace(user.Status)

	if isPartial {
		if user.Firstname != "" {
			savedUser.Firstname = user.Firstname
		}

		if user.Lastname != "" {
			savedUser.Lastname = user.Lastname
		}

		if user.Email != "" {
			savedUser.Email = user.Email
		}

		if user.Status != "" {
			savedUser.Status = user.Status
		}
	} else {
		savedUser.Firstname = user.Firstname
		savedUser.Lastname = user.Lastname
		savedUser.Email = user.Email
		savedUser.Status = user.Status
	}

	err = savedUser.Update()
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func DeleteUser(userId int64) *errors.RestError {
	return users.DeleteById(userId)
}

func GetUserByStatus(status string) ([]*users.User, *errors.RestError) {
	return users.GetUserByStatus(status)
}