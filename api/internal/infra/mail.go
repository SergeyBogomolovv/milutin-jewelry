package infra

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"gopkg.in/gomail.v2"
)

type mailService struct {
	log  *slog.Logger
	host string
	port int
	user string
	pass string
	to   string
}

func NewMailService(log *slog.Logger, cfg config.MailConfig, to string) *mailService {
	return &mailService{
		log:  log,
		host: cfg.Host,
		port: cfg.Port,
		user: cfg.User,
		pass: cfg.Pass,
		to:   to,
	}
}

func (s *mailService) SendCodeToAdmin(ctx context.Context, code string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", s.to)

	m.SetHeader("Subject", "Вход в админ панель milutin-jewelry")
	m.SetBody("text/html", fmt.Sprintf("Код авторизации: <b>%s</b>", code))

	d := gomail.NewDialer(s.host, s.port, s.user, s.pass)

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		errChan <- d.DialAndSend(m)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
