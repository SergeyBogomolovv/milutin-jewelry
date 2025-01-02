package mail

import (
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

func New(log *slog.Logger, cfg config.MailConfig, to string) *mailService {
	const dest = "mailService"
	return &mailService{
		log:  log.With(slog.String("dest", dest)),
		host: cfg.Host,
		port: cfg.Port,
		user: cfg.User,
		pass: cfg.Pass,
		to:   to,
	}
}

func (s *mailService) SendCodeToAdmin(code string) {
	const op = "SendCodeToAdmin"
	log := s.log.With(slog.String("op", op))

	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", s.to)

	m.SetHeader("Subject", "Вход в админ панель milutin-jewelry")
	m.SetBody("text/html", messageBody(code))

	d := gomail.NewDialer(s.host, s.port, s.user, s.pass)

	if err := d.DialAndSend(m); err != nil {
		log.Error("failed to send email", "err", err)
		return
	}
	log.Info("email sent")
}