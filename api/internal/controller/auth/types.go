package auth

import (
	"context"
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
)

type Usecase interface {
	Login(ctx context.Context, code string) (string, error)
	SendCode(ctx context.Context) error
}

type LoginBody struct {
	Code string `json:"code" validate:"len=6"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func WriteToken(w http.ResponseWriter, token string, status int) error {
	return res.WriteJSON(w, TokenResponse{Token: token}, status)
}
