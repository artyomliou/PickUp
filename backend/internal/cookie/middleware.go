package cookie

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Print("Recovered. Error:", r)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  "error",
					"message": "必須先登入才能呼叫這個API",
				})
			}
		}()

		cookie, err := c.Cookie(CookieName)
		if cookie == "" || err != nil {
			panic("Cookie is empty")
		}
		token := ParseToken(cookie)
		if !token.Valid {
			panic("Token is not valid.")
		}
		c.Set("verified-token", token)
		c.Next()
	}
}
