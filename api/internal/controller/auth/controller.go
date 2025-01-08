package auth

import (
	"errors"
	"fmt"
	"net/http"

	usecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	validate *validator.Validate
	usecase  Usecase
}

func Register(router *http.ServeMux, usecase Usecase) {
	const dest = "authController"
	controller := &controller{
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
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
// @Failure      400    {object}  res.ErrorResponse  "Неверный запрос"
// @Failure      403    {object}  res.ErrorResponse  "Неверный код"
// @Failure      500    {object}  res.ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /auth/login [post]
func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginBody
	if err := res.DecodeBody(r, &payload); err != nil {
		res.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	if err := c.validate.Struct(payload); err != nil {
		res.WriteError(w, fmt.Sprintf("invalid payload: %v", err), http.StatusBadRequest)
		return
	}

	token, err := c.usecase.Login(r.Context(), payload.Code)

	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCode) {
			res.WriteError(w, "invalid code", http.StatusForbidden)
			return
		}
		res.WriteError(w, "failed to check login code", http.StatusInternalServerError)
		return
	}

	WriteToken(w, token, http.StatusOK)
}

// @Summary      Отправка кода
// @Description  Отправляет код на почту админа
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      201    {object}  res.MessageResponse
// @Failure      500    {object}  res.ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /auth/send-code [post]
func (c *controller) SendVerificationCode(w http.ResponseWriter, r *http.Request) {
	if err := c.usecase.SendCode(r.Context()); err != nil {
		res.WriteError(w, "failed to send code", http.StatusInternalServerError)
		return
	}
	res.WriteMessage(w, "code sent", http.StatusCreated)
}
