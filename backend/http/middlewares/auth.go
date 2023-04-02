package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const CookieName = "auth"

var KeyFunc = func(*jwt.Token) (interface{}, error) {
	return "HS256", nil
}

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

// TODO no middleware impl
