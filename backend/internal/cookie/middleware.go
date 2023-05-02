package cookie

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(CookieName)
		if cookie == "" || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "必須先登入才能呼叫這個API",
			})
			return
		}

		token := ParseToken(cookie)
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token 不正確 #1",
			})
			return
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token 不正確 #2",
			})
			return
		}

		uid, err := strconv.Atoi(sub)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token 不正確 #3",
			})
			return
		}

		c.Set("uid", uid)
		c.Next()
	}
}
