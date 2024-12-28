package middleware

import (
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

func NewAuthMiddleware(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) < 8 {
				utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			token := authHeader[len("Bearer "):]
			if token == "" {
				utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			if err := utils.VerifyToken(token, []byte(secret)); err != nil {
				utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
