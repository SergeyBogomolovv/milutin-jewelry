package auth

import (
	"context"
	"crypto/subtle"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/code"
)

type usecase struct {
	log       *slog.Logger
	codes     CodeStorage
	mail      MailService
	admin     config.AdminConfig
	jwtSecret []byte
	jwtTTL    time.Duration
}

func New(log *slog.Logger, codes CodeStorage, mail MailService, admin config.AdminConfig, conf config.JwtConfig) *usecase {
	const dest = "authUsecase"
	return &usecase{
		log:       log.With(slog.String("dest", dest)),
		codes:     codes,
		mail:      mail,
		admin:     admin,
		jwtSecret: conf.Secret,
		jwtTTL:    conf.TTL,
	}
}

func (u *usecase) SendCode(ctx context.Context) error {
	const op = "SendCode"
	log := u.log.With(slog.String("op", op))

	code, err := u.codes.Create(ctx)
	if err != nil {
		log.Error("can't create login code", "err", err)
		return err
	}

	if err := u.mail.SendCodeToAdmin(code); err != nil {
		log.Error("can't send login code to admin", "err", err)
		return err
	}

	return nil
}

func (u *usecase) Login(ctx context.Context, code string) (string, error) {
	const op = "Login"
	log := u.log.With(slog.String("op", op))

	if err := u.codes.Check(ctx, code); err != nil {
		if errors.Is(err, storage.ErrInvalidCode) {
			log.Info("invalid login code")
			return "", ErrInvalidCode
		}
		log.Error("can't check login code", "err", err)
		return "", err
	}

	if err := u.codes.Delete(ctx, code); err != nil {
		log.Error("can't delete login code", "err", err)
		return "", err
	}

	token, err := u.signToken()
	if err != nil {
		log.Error("can't sign token", "err", err)
		return "", err
	}

	return token, nil
}

func (u *usecase) LoginByPassword(ctx context.Context, email, password string) (string, error) {
	const op = "LoginByPassword"
	log := u.log.With(slog.String("op", op))

	if !u.checkCredentials(email, password) {
		log.Info("invalid admin credentials")
		return "", ErrInvalidCredentials
	}

	token, err := u.signToken()
	if err != nil {
		log.Error("can't sign token", "err", err)
		return "", err
	}

	return token, nil
}

func (u *usecase) checkCredentials(email, password string) bool {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	emailMatch := subtle.ConstantTimeCompare([]byte(email), []byte(u.admin.Email)) == 1
	passwordMatch := subtle.ConstantTimeCompare([]byte(password), []byte(u.admin.Password)) == 1
	return emailMatch && passwordMatch
}
