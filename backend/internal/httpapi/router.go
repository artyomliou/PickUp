package httpapi

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(corsMiddleware())

	// Routes
	api := router.Group("/api")
	{
		loginRoutes(api)
	}

	return
}

func loginRoutes(rg *gin.RouterGroup) {
	login := new(LoginController)
	rg.GET("/is-logged-in", login.IsLoggedIn)
	rg.POST("/login", login.Login)
	rg.GET("/logout", login.Logout)
}
