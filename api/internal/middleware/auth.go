package middleware

import (
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

type Middleware func(http.Handler) http.Handler

func NewAuthMiddleware(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			tokenString := cookie.Value
			if err := utils.VerifyToken(tokenString, []byte(secret)); err != nil {
				utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
