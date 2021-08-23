package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/golang-jwt/jwt"
)

type JWTMiddleware struct {
	UsersService interfaces.IUsersService
}

type ContextKey string

const ContextUserKey ContextKey = "user"

var signingKey []byte

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	signingKey = []byte(jwtKey)
}

func parseJwtToken(tokenString string) (userId int, err error) {

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

// Middleware function, which will be called for each request
func (m *JWTMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		userId, err := parseJwtToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := m.UsersService.Get(userId)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origins := os.Getenv("ORIGINS")
		w.Header().Set("Access-Control-Allow-Origin", origins)
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Ok.")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
