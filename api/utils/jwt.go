package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userIdAsString string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims =
		&jwt.StandardClaims{
			Subject:   userIdAsString,
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		}

	return t.SignedString(os.Getenv("JWT_KEY"))
}
