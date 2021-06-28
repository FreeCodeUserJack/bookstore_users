package mysql_utils

import (
	"strings"

	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	NoUserFoundForId = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), NoUserFoundForId) {
			return errors.NewNotFoundError("no record matching given id", err.Error())
		}
		return errors.NewInternalServerError("could not convert to MySQLError", err.Error())
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated key", sqlErr.Error())
	default:
		return errors.NewInternalServerError("MySQLError occurred", sqlErr.Error())
	}
}