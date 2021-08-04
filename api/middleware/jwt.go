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
		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]

		userId, err := utils.ParseToken(tokenString)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := m.UserModel.GetById(userId)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
