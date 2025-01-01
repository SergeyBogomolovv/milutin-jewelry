package authusecase

import (
	"context"
	"errors"
)

type CodeStorage interface {
	Check(ctx context.Context, code string) error
	Create(ctx context.Context) (string, error)
	Delete(ctx context.Context, code string) error
}

type MailService interface {
	SendCodeToAdmin(code string)
}

var (
	ErrInvalidCode = errors.New("invalid otp code")
)
