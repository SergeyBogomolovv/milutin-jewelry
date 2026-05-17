package middleware

import (
	"net/http"
	"strings"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

func NewAuthMiddleware(secret []byte) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token, ok := strings.CutPrefix(authHeader, "Bearer ")
			if !ok {
				res.WriteError(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			token = strings.TrimSpace(token)
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
