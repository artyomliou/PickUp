package cookie

import (
	"errors"
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

		uid, err := ParseTokenForUid(cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "資料異常，請重新登入",
			})
			return
		}

		c.Set("uid", uid)
		c.Next()
	}
}

func ParseTokenForUid(tokenStr string) (uint, error) {
	token, err := ParseToken(tokenStr)
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("Token is invalid")
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return 0, errors.New("Cannot get subject of token")
	}

	uid, err := strconv.Atoi(sub)
	if err != nil {
		return 0, errors.New("Cannot convert subject from string to uint")
	}

	return uint(uid), nil
}
