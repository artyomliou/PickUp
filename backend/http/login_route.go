package http

import "github.com/gin-gonic/gin"

func loginRoutes(rg *gin.RouterGroup) {
	login := new(LoginController)
	rg.GET("/is-logged-in", login.IsLoggedIn)
	rg.POST("/login", login.Login)
	rg.GET("/logout", login.Logout)
}
