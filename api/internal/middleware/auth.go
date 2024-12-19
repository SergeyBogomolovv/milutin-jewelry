package middleware

import (
	"fmt"
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

type Middleware func(http.Handler) http.Handler

func NewAuthMiddleware(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")[len("Bearer "):]
			fmt.Println(token)
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
