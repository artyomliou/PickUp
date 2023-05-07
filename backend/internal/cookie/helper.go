package cookie

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeCookieString(token string) string {
	return fmt.Sprintf("%s=%s", CookieName, token)
}

func CreateToken(userId uint, expiredAt time.Time) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userId),
		ID:        uuid.NewString(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if ss, err := token.SignedString(appKeyBytes); err != nil {
		return "", err
	} else {
		return ss, nil
	}
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return appKeyBytes, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
