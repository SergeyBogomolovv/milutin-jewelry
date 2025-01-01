package authcontroller

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	usecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	validate *validator.Validate
	usecase  AuthUsecase
	log      *slog.Logger
}

func Register(log *slog.Logger, router *http.ServeMux, usecase AuthUsecase) {
	const dest = "authController"
	controller := &authController{
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
		log:      log.With(slog.String("dest", dest)),
	}
	r := http.NewServeMux()
	r.HandleFunc("POST /login", controller.Login)
	r.HandleFunc("POST /send-code", controller.SendVerificationCode)

	router.Handle("/auth/", http.StripPrefix("/auth", r))
}

// @Summary      Вход
// @Description  Нужен код с почты админа, отправляет jwt токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      LoginBody  true  "Данные для входа"
// @Success      200    {object}  TokenResponse
// @Failure      400    {object}  ErrorResponse  "Неверный запрос"
// @Failure      403    {object}  ErrorResponse  "Неверный код"
// @Failure      500    {object}  ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /auth/login [post]
func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	const op = "Login"
	log := c.log.With(slog.String("op", op))

	var payload LoginBody
	if err := utils.DecodeBody(r, &payload); err != nil {
		log.Error("failed to decode payload", "err", err)
		WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	if err := c.validate.Struct(payload); err != nil {
		WriteError(w, fmt.Sprintf("invalid payload: %v", err), http.StatusBadRequest)
		return
	}

	token, err := c.usecase.Login(r.Context(), payload.Code)

	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCode) {
			WriteError(w, "invalid code", http.StatusForbidden)
			return
		}
		WriteError(w, "failed to check login code", http.StatusInternalServerError)
		return
	}

	WriteToken(w, token, http.StatusOK)
}

// @Summary      Отправка кода
// @Description  Отправляет код на почту админа
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200    {object}  MessageResponse
// @Failure      500    {object}  ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /auth/send-code [post]
func (c *authController) SendVerificationCode(w http.ResponseWriter, r *http.Request) {
	if err := c.usecase.SendCode(r.Context()); err != nil {
		utils.WriteError(w, "failed to send code", http.StatusInternalServerError)
		return
	}
	WriteMessage(w, "code sent", http.StatusCreated)
}
