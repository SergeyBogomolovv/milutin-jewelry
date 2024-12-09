package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthUsecase interface{}

type authController struct {
	validate validator.Validate
	uc       AuthUsecase
}

func NewAuthController(uc AuthUsecase) *authController {
	return &authController{
		uc:       uc,
		validate: *validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {}

func (c *authController) SendCode(w http.ResponseWriter, r *http.Request) {}
