package authcontroller

import (
	"context"
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
)

type AuthUsecase interface {
	Login(ctx context.Context, code string) (string, error)
	SendCode(ctx context.Context) error
}

type LoginBody struct {
	Code string `json:"code"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func WriteToken(w http.ResponseWriter, token string, status int) error {
	return utils.WriteJSON(w, TokenResponse{Token: token}, status)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, err string, status int) {
	utils.WriteJSON(w, ErrorResponse{Error: err}, status)
}

type MessageResponse struct {
	Message string `json:"message"`
}

func WriteMessage(w http.ResponseWriter, msg string, status int) {
	utils.WriteJSON(w, MessageResponse{Message: msg}, status)
}
