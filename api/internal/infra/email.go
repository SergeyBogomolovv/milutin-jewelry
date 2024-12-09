package infra

import (
	"context"
	"fmt"
	"net/smtp"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
)

type smtpEmailSender struct {
	host       string
	port       string
	username   string
	adminEmail string
	auth       smtp.Auth
}

func NewEmailSender(cfg config.MailConfig, adminEmail string) *smtpEmailSender {
	return &smtpEmailSender{
		host:       cfg.Host,
		port:       cfg.Port,
		username:   cfg.User,
		adminEmail: adminEmail,
		auth:       smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host),
	}
}

func (s *smtpEmailSender) SendCodeToAdmin(ctx context.Context, code string) error {
	sendCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", s.adminEmail, "Код авторизации", code)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	errCh := make(chan error, 1)
	go func() {
		errCh <- smtp.SendMail(addr, s.auth, s.username, []string{s.adminEmail}, []byte(msg))
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-sendCtx.Done():
		return fmt.Errorf("email send operation timed out: %w", sendCtx.Err())
	}

	return nil
}
