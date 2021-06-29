package users

import (
	"strings"

	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status 			string `json:"status"`
	Password 		string `json:"-"`
}

func (u *User) Validate() *errors.RestError {
	u.Firstname = strings.TrimSpace(u.Firstname)
	u.Lastname = strings.TrimSpace(u.Lastname)

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address", "bad user request")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("password is not valid", "invalid passowrd")
	}

	return nil
}