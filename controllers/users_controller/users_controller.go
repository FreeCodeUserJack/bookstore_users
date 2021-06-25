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
	userIdInt, err := strconv.ParseInt(strings.TrimSpace(c.Param("userId")), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("userId must be a valid number", "bad_request")
		c.JSON(restErr.Status, restErr)
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
	c.String(http.StatusNotImplemented, "implement me")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}