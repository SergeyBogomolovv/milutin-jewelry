package middleware

import (
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

func NewAuthMiddleware(secret []byte) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) < 8 {
				res.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			token := authHeader[len("Bearer "):]
			if token == "" {
				res.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			if err := utils.VerifyToken(token, secret); err != nil {
				res.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
