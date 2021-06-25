package users

import (
	"fmt"

	"github.com/FreeCodeUserJack/bookstore_users/util/date_utils"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
)


var (
	usersDB = make(map[int64]*User)
)

func GetUserById(userId int64) (*User, *errors.RestError) {
	res, ok := usersDB[userId]
	if !ok {
		return nil, errors.NewNotFoundError("invalid user ID", "user not found in DB")
	}
	
	return res, nil
}

func (u *User) Save() *errors.RestError {
	if user := usersDB[u.Id]; user != nil {
		if user.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user with email %s already registered", u.Email), "bad request")
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with id %d already exists", u.Id), "bad request")
	}

	u.DateCreated = date_utils.GetTimeNowString()

	usersDB[u.Id] = u
	return nil
}