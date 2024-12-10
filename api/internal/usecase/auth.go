package usecase

import (
	"context"
	"log/slog"
	"time"

	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/golang-jwt/jwt/v5"
)

type CodesRepo interface {
	Check(ctx context.Context, code string) error
	Create(ctx context.Context) (string, error)
	Delete(ctx context.Context, code string) error
}

type MailService interface {
	SendCodeToAdmin(ctx context.Context, code string) error
}

type authUsecase struct {
	log       *slog.Logger
	cr        CodesRepo
	ms        MailService
	jwtSecret []byte
}

func NewAuthUsecase(log *slog.Logger, cr CodesRepo, ms MailService, jwtSecret string) *authUsecase {
	return &authUsecase{log: log, cr: cr, ms: ms, jwtSecret: []byte(jwtSecret)}
}

func (u *authUsecase) SendCode(ctx context.Context) error {
	log := u.log.With(slog.String("op", "SendCode"))

	log.Info("sending login code")

	code, err := u.cr.Create(ctx)
	if err != nil {
		log.Error("failed to create login code", "err", err)
		return err
	}

	if err := u.ms.SendCodeToAdmin(ctx, code); err != nil {
		log.Error("failed to send email", "err", err)
		return err
	}

	log.Info("login code sent")
	return nil
}

func (u *authUsecase) Login(ctx context.Context, code string) (string, error) {
	log := u.log.With(slog.String("code", code), slog.String("op", "Login"))

	log.Info("logging in")

	if err := u.cr.Check(ctx, code); err != nil {
		log.Info("invalid login code")
		return "", errs.ErrInvalidLoginCode
	}

	if err := u.cr.Delete(ctx, code); err != nil {
		log.Error("failed to delete login code", "err", err)
		return "", err
	}

	token, err := u.signToken()
	if err != nil {
		log.Error("failed to sign token", "err", err)
		return "", err
	}

	log.Info("logged in")

	return token, nil
}

func (u *authUsecase) signToken() (string, error) {
	iat := time.Now()
	exp := iat.Add(24 * time.Hour)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(iat),
		ExpiresAt: jwt.NewNumericDate(exp),
	}).SignedString(u.jwtSecret)
}
