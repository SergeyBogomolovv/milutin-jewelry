package auth

import (
	"context"
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
)

type Usecase interface {
	Login(ctx context.Context, code string) (string, error)
	LoginByPassword(ctx context.Context, email, password string) (string, error)
	SendCode(ctx context.Context) error
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func WriteToken(w http.ResponseWriter, token string, status int) error {
	return res.WriteJSON(w, TokenResponse{Token: token}, status)
}
