package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addLoginRoutes(rg *gin.RouterGroup) {
	rg.GET("/is-logged-in", func(c *gin.Context) {
		cookie, err := c.Cookie("vp")
		c.JSON(http.StatusOK, gin.H{
			"status":       "success",
			"is-logged-in": len(cookie) > 0 && err == nil,
		})
	})
	rg.GET("/login", func(c *gin.Context) {
		token := "JWTJWTJWT"
		maxAge := 86400 * 7
		c.SetCookie("vp", token, maxAge, "/api", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})
	})
	rg.GET("/logout", func(c *gin.Context) {
		token := ""
		maxAge := 0
		c.SetCookie("vp", token, maxAge, "/api", "localhost", false, true)
		c.JSON(http.StatusNoContent, nil)
	})
}
