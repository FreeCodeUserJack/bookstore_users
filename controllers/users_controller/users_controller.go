package users_controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/FreeCodeUserJack/bookstore_users/domain/users"
	"github.com/FreeCodeUserJack/bookstore_users/services"
	"github.com/FreeCodeUserJack/bookstore_users/util/errors"
	"github.com/gin-gonic/gin"
)


func GetUser(c *gin.Context) {
	userIdInt, err := getUserId(c.Param("userId"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	foundUser, restErr := services.GetUserById(userIdInt)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, foundUser)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	res, restErr := services.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := getUserId(c.Param("userId"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	updatedUser, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userId, err := getUserId(c.Param("userId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	delErr := services.DeleteUser(userId)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"deleted user with id": c.Param("userId")})
}

func getUserId(userId string) (int64, *errors.RestError) {
	userIdInt, err := strconv.ParseInt(strings.TrimSpace(userId), 10, 64)
	if err != nil {
		return -1, errors.NewBadRequestError("invalid user id", err.Error())
	}

	return userIdInt, nil
}

func GetUserByStatus(c *gin.Context) {
	status := strings.TrimSpace(c.Param("status"))
	if status == "" {
		err := errors.NewBadRequestError("status must be give, cannot be empty", "bad request")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	users, err := services.GetUserByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
}