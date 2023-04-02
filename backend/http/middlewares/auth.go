package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const CookieName = "auth"

func CreateToken(userId uint, expiredAt time.Time) (ss string) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userId),
		ID:        uuid.NewString(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	if ss, err = token.SignedString([]byte(os.Getenv("APP_KEY"))); err != nil {
		log.Panic(err)
	}
	return ss
}

func ParseToken(tokenStr string) *jwt.Token {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("APP_KEY")), nil
	})
	if err != nil {
		log.Panic(err)
	}
	return token
}

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
