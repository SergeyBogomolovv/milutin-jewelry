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
	SendCodeToAdmin(code string)
}

type authUsecase struct {
	log       *slog.Logger
	cr        CodesRepo
	ms        MailService
	jwtSecret []byte
}

func NewAuthUsecase(log *slog.Logger, cr CodesRepo, ms MailService, jwtSecret string) *authUsecase {
	return &authUsecase{log: log.With(slog.String("op", "authUsecase")), cr: cr, ms: ms, jwtSecret: []byte(jwtSecret)}
}

func (u *authUsecase) SendCode(ctx context.Context) error {
	code, err := u.cr.Create(ctx)
	if err != nil {
		u.log.Error("failed to create login code", "err", err)
		return err
	}

	go u.ms.SendCodeToAdmin(code)

	return nil
}

func (u *authUsecase) Login(ctx context.Context, code string) (string, error) {
	if err := u.cr.Check(ctx, code); err != nil {
		u.log.Info("invalid login code")
		return "", errs.ErrInvalidLoginCode
	}

	if err := u.cr.Delete(ctx, code); err != nil {
		u.log.Error("failed to delete login code", "err", err)
		return "", err
	}

	token, err := u.signToken()
	if err != nil {
		u.log.Error("failed to sign token", "err", err)
		return "", err
	}

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
