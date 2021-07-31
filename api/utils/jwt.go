package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userIdAsString string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims =
		&jwt.StandardClaims{
			Subject:   userIdAsString,
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		}

	jwtKey := os.Getenv("JWT_KEY")

	return token.SignedString([]byte(jwtKey))
}
