package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()

	api := router.Group("/api")
	{
		loginRoutes(api)
	}

	return
}
