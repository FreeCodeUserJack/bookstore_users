package app

import (
	"github.com/FreeCodeUserJack/bookstore_users/controllers/ping_controller"
	"github.com/FreeCodeUserJack/bookstore_users/controllers/users_controller"
)


func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:userId", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
	router.PUT("/users/:userId", users_controller.UpdateUser)
	router.PATCH("/users/:userId", users_controller.UpdateUser)
	router.DELETE("/users/:userId", users_controller.DeleteUser)
	router.GET("/internal/users/search", users_controller.Search)
}