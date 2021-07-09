package services

import (
	"strings"

	"github.com/FreeCodeUserJack/bookstore_users/domain/users"
	"github.com/FreeCodeUserJack/bookstore_users/util/crypto_utils"
	"github.com/FreeCodeUserJack/bookstore_users/util/date_utils"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUserById(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetDbTimeNowString()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUserById(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("userId cannot be <= 0", "bad_request")
	}

	return users.GetUserById(userId)
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	savedUser, err := UsersService.GetUserById(user.Id)
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

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	return users.DeleteById(userId)
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	return users.GetUserByStatus(status)
}