package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Christheoreo/project-manager/models"
	"github.com/Christheoreo/project-manager/utils"
)

type JWTMiddleware struct {
	UserModel models.User
}

type ContextKey string

const ContextUserKey ContextKey = "user"

// Middleware function, which will be called for each request
func (m *JWTMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenArray := strings.Split(tokenString, " ")
		tokenString = tokenArray[1]

		userId, err := utils.ParseToken(tokenString)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		user, err := m.UserModel.GetById(userId)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
