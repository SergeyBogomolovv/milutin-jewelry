package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type AuthUsecase interface {
	Login(ctx context.Context, code string) (string, error)
	SendCode(ctx context.Context) error
}

type authController struct {
	validate validator.Validate
	uc       AuthUsecase
}

func RegisterAuthController(router *http.ServeMux, uc AuthUsecase) {
	controller := &authController{
		uc:       uc,
		validate: *validator.New(validator.WithRequiredStructEnabled()),
	}
	router.HandleFunc("POST /login", controller.Login)
	router.HandleFunc("POST /send_login_code", controller.SendCode)
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginDTO
	if err := utils.DecodeBody(r, &payload); err != nil {
		utils.WriteError(w, "failed to decode payload", http.StatusBadRequest)
		return
	}
	if err := c.validate.Struct(payload); err != nil {
		utils.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}

	token, err := c.uc.Login(r.Context(), payload.Code)

	if err != nil {
		if errors.Is(err, errs.ErrInvalidLoginCode) {
			utils.WriteError(w, "invalid code", http.StatusBadRequest)
			return
		}
		utils.WriteError(w, "failed to check login code", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	utils.WriteMessage(w, "login successful", http.StatusOK)
}

func (c *authController) SendCode(w http.ResponseWriter, r *http.Request) {
	if err := c.uc.SendCode(r.Context()); err != nil {
		utils.WriteError(w, "failed to send code", http.StatusInternalServerError)
		return
	}
	utils.WriteMessage(w, "code sent", http.StatusCreated)
}
