package users

import (
	"fmt"
	"strings"

	"github.com/FreeCodeUserJack/bookstore_users/datasources/mysql/users_db"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
	"github.com/FreeCodeUserJack/bookstore_users/util/mysql_utils"
	"github.com/FreeCodeUserJack/bookstore_utils/logger"
	// "github.com/FreeCodeUserJack/bookstore_users/util/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryDeleteUserById = "DELETE FROM users WHERE id=?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func GetUserById(userId int64) (*User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("could not prepare query statement", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	defer stmt.Close()

	sqlRow := stmt.QueryRow(userId)

	resUser := &User{}

	err = sqlRow.Scan(&resUser.Id, &resUser.Firstname, &resUser.Lastname, &resUser.Email, &resUser.DateCreated, &resUser.Status)
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	
	return resUser, nil
}

func (u *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("could not prepare save user query statement", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(u.Firstname, u.Lastname, u.Email, u.DateCreated, u.Status, u.Password)
	if saveErr != nil {
		logger.Error("error when trying to Exec save user stmt", saveErr)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get insert id for save user query", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	u.Id = userId
	return nil
}

func (u *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare stmt for update user query", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Firstname, u.Lastname, u.Email, u.Status, u.Id)
	if err != nil {
		logger.Error("error when trying to Exec update user query", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	return nil
}

func DeleteById(userId int64) *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUserById)
	if err != nil {
		logger.Error("error when trying to Prepare delete user query", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	res, delErr := stmt.Exec(userId)
	if delErr != nil {
		logger.Error("error when trying to Exec delete user query", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	rows, err := res.RowsAffected()

	if err != nil {
		logger.Error("error when trying to get RowsAffected for delete user", err)
		return errors.NewInternalServerError("database error", "internal server error")
	}

	if rows == 0 {
		logger.Error("error b/c deleteById did not find the user, bad request", err)
		return errors.NewBadRequestError("user id not found in DB", "bad request")
	}

	return nil
}

func GetUserByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to Prepare stmt for SearchUser", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to Query() for SearchUser", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	defer rows.Close()

	usersRes := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			logger.Error("error when trying to Scan row into User{} for SearchUser", err)
			return nil, errors.NewInternalServerError("database error", "internal server error")
		}
		usersRes = append(usersRes, user)
	}

	if len(usersRes) == 0 {
		// don't log b/c if someone is using API wrong, there would be a TON of logs, not good
		// logger.Error("error b/c SearchUser did not find any user with corresponding status " + status, err)
		return nil, errors.NewNotFoundError(fmt.Sprintf("status: %s not found", status), "no users found")
	}

	return usersRes, nil
}

func FindByEmailAndPassword(email, password string) (*User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("could not prepare query statement for find by email/password", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	defer stmt.Close()

	sqlRow := stmt.QueryRow(email, password, StatusActive)

	resUser := &User{}

	err = sqlRow.Scan(&resUser.Id, &resUser.Firstname, &resUser.Lastname, &resUser.Email, &resUser.DateCreated, &resUser.Status)
	if err != nil {
		if strings.Contains(err.Error(), mysql_utils.NoUserFoundForId) {
			return nil, errors.NewBadRequestError("no user found with given credentials", "bad request")
		}
		logger.Error("error when trying to get user by email and password", err)
		return nil, errors.NewInternalServerError("database error", "internal server error")
	}
	
	return resUser, nil
}