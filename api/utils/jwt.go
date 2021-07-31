package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var signingKey []byte

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	signingKey = []byte(jwtKey)
}

func CreateToken(userIdAsString string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims =
		&jwt.StandardClaims{
			Subject:   userIdAsString,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		}

	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) (userId primitive.ObjectID, err error) {

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
		userId, err = primitive.ObjectIDFromHex(userIdString)
	}

	return
}
