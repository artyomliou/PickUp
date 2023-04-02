package http

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()

	if os.Getenv("APP_DEBUG") == "1" {
		router.StaticFile("/debug", "./public/debug.html")
	}

	api := router.Group("/api")
	{
		loginRoutes(api)
	}

	return
}
