package app

import (
	"github.com/FreeCodeUserJack/bookstore_utils/logger"
	"github.com/gin-gonic/gin"
)


var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()

	logger.Info("about to start application")

	router.Run(":8000")
}