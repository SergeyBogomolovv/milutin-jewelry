package auth

import (
	"context"
	"errors"
	"log/slog"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/code"
)

type usecase struct {
	log       *slog.Logger
	codes     CodeStorage
	mail      MailService
	jwtSecret []byte
}

func New(log *slog.Logger, codes CodeStorage, mail MailService, jwtSecret string) *usecase {
	const dest = "authUsecase"
	return &usecase{
		log:       log.With(slog.String("dest", dest)),
		codes:     codes,
		mail:      mail,
		jwtSecret: []byte(jwtSecret),
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

	go u.mail.SendCodeToAdmin(code)

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
