package infra

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
)

type mailService struct {
	log  *slog.Logger
	host string
	port string
	user string
	pass string
	to   string
	tls  *tls.Config
}

func NewMailService(log *slog.Logger, cfg config.MailConfig, to string) *mailService {
	return &mailService{
		log:  log,
		host: cfg.Host,
		port: cfg.Port,
		user: cfg.User,
		pass: cfg.Pass,
		to:   to,
		tls:  &tls.Config{InsecureSkipVerify: false, ServerName: cfg.Host},
	}
}

func (s *mailService) SendCodeToAdmin(ctx context.Context, code string) error {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port), s.tls)
	if err != nil {
		s.log.Error("failed to establish tls connection", "err", err)
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.host)
	if err != nil {
		s.log.Error("failed to connect to smtp server", "err", err)
		return err
	}
	defer client.Close()

	if err := client.Auth(smtp.PlainAuth("", s.user, s.pass, s.host)); err != nil {
		s.log.Error("failed to authenticate", "err", err)
		return err
	}

	if err := client.Mail(s.user); err != nil {
		s.log.Error("failed to add sender", "err", err)
		return err
	}

	if err := client.Rcpt(s.to); err != nil {
		s.log.Error("failed to add recipient", "err", err)
		return err
	}

	wc, err := client.Data()
	if err != nil {
		s.log.Error("failed to open data stream", "err", err)
		return err
	}

	_, err = wc.Write(createCodeEmail(code))
	if err != nil {
		s.log.Error("failed to write data", "err", err)
		return err
	}

	if err := wc.Close(); err != nil {
		s.log.Error("failed to close data stream", "err", err)
		return err
	}

	return client.Quit()
}

func createCodeEmail(code string) []byte {
	var buff bytes.Buffer
	buff.WriteString("Subject: Вход в админ панель milutin-jewelry\r\n")
	buff.WriteString("MIME-Version: 1.0\r\n")
	buff.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
	buff.WriteString(fmt.Sprintf("<p>Код авторизации: <strong>%s</strong></p>", code))
	return buff.Bytes()
}
