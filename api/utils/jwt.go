package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey []byte

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	signingKey = []byte(jwtKey)
}

func CreateToken(userId int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims =
		&jwt.StandardClaims{
			Subject:   strconv.Itoa(userId),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		}

	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) (userId int, err error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdString := claims["sub"].(string)
		userId, err = strconv.Atoi(userIdString)
	}
	return
}
