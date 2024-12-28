package infra

import (
	"fmt"
	"log/slog"

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

func (s *mailService) SendCodeToAdmin(code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", s.to)

	m.SetHeader("Subject", "Вход в админ панель milutin-jewelry")
	m.SetBody("text/html", fmt.Sprintf("Код авторизации: <b>%s</b>. Код действителен в течении 5 минут", code))

	d := gomail.NewDialer(s.host, s.port, s.user, s.pass)

	if err := d.DialAndSend(m); err != nil {
		s.log.Error("failed to send email", "err", err)
		return
	}
	s.log.Info("email sent")
}
