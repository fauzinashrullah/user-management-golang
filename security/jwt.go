package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("secret_test")

func GenerateJwt(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
