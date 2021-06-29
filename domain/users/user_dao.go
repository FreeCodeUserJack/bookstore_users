package users

import (
	"fmt"

	"github.com/FreeCodeUserJack/bookstore_users/datasources/mysql/users_db"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
	"github.com/FreeCodeUserJack/bookstore_users/util/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryDeleteUserById = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func GetUserById(userId int64) (*User, *errors.RestError) {
	if err := users_db.Client.Ping(); err != nil {
		return nil, errors.NewInternalServerError("db refused connection", "server error")
	}

	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return nil, errors.NewInternalServerError("could not prepare query statement", err.Error())
	}
	defer stmt.Close()

	sqlRow := stmt.QueryRow(userId)

	resUser := &User{}

	err = sqlRow.Scan(&resUser.Id, &resUser.Firstname, &resUser.Lastname, &resUser.Email, &resUser.DateCreated, &resUser.Status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	
	return resUser, nil
}

func (u *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("could not prepare query statement", err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(u.Firstname, u.Lastname, u.Email, u.DateCreated, u.Status, u.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	u.Id = userId
	return nil
}

func (u *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Firstname, u.Lastname, u.Email, u.Status, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func DeleteById(userId int64) *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUserById)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	res, delErr := stmt.Exec(userId)
	if delErr != nil {
		return mysql_utils.ParseError(delErr)
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return errors.NewInternalServerError("rows affected func not supported", err.Error())
	}

	if rows == 0 {
		return errors.NewBadRequestError(fmt.Sprintf("userid: %d not in DB", userId), "user id not found")
	}

	return nil
}

func GetUserByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	usersRes := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		usersRes = append(usersRes, user)
	}

	if len(usersRes) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("status: %s not found", status), "no users found")
	}

	return usersRes, nil
}