package mail

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/mail"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"gopkg.in/gomail.v2"
)

const (
	senderName     = "Milutin Jewelry"
	sendAttempts   = 3
	sendRetryDelay = 500 * time.Millisecond
	dialTimeout    = 5 * time.Second
)

type mailService struct {
	log  *slog.Logger
	host string
	port int
	user string
	pass string
	to   string
}

func New(log *slog.Logger, cfg config.MailConfig) *mailService {
	const dest = "mailService"
	return &mailService{
		log:  log.With(slog.String("dest", dest)),
		host: cfg.Host,
		port: cfg.Port,
		user: cfg.User,
		pass: cfg.Pass,
		to:   cfg.Admin,
	}
}

func (s *mailService) SendCodeToAdmin(code string) error {
	const op = "SendCodeToAdmin"
	log := s.log.With(slog.String("op", op))

	if err := s.validate(); err != nil {
		log.Error("invalid mail config", "err", err)
		return err
	}

	var errs []error
	for attempt := 1; attempt <= sendAttempts; attempt++ {
		if err := s.send(code); err != nil {
			errs = append(errs, err)
			log.Warn("failed to send email", "attempt", attempt, "host", s.host, "port", s.port, "err", err)

			if isNetworkTimeout(err) {
				break
			}

			if attempt < sendAttempts {
				time.Sleep(time.Duration(attempt) * sendRetryDelay)
			}
			continue
		}

		log.Info("email sent", "attempt", attempt)
		return nil
	}

	return fmt.Errorf("failed to send email after %d attempts: %w", sendAttempts, errors.Join(errs...))
}

func (s *mailService) send(code string) error {
	if err := s.checkSMTPReachable(); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", s.user, senderName)
	m.SetAddressHeader("To", s.to, "")
	m.SetHeader("Subject", "Вход в админ панель milutin-jewelry")
	m.SetBody("text/plain", plainMessageBody(code))
	m.AddAlternative("text/html", htmlMessageBody(code))

	d := gomail.NewDialer(s.host, s.port, s.user, s.pass)
	return d.DialAndSend(m)
}

func (s *mailService) checkSMTPReachable() error {
	address := net.JoinHostPort(s.host, fmt.Sprint(s.port))
	conn, err := net.DialTimeout("tcp", address, dialTimeout)
	if err != nil {
		return fmt.Errorf("smtp endpoint %s is unreachable: %w", address, err)
	}
	return conn.Close()
}

func isNetworkTimeout(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func (s *mailService) validate() error {
	if s.host == "" {
		return errors.New("mail host is empty")
	}
	if s.port <= 0 {
		return fmt.Errorf("invalid mail port: %d", s.port)
	}
	if s.user == "" {
		return errors.New("mail user is empty")
	}
	if s.pass == "" {
		return errors.New("mail password is empty")
	}
	if _, err := mail.ParseAddress(s.user); err != nil {
		return fmt.Errorf("invalid sender email: %w", err)
	}
	if _, err := mail.ParseAddress(s.to); err != nil {
		return fmt.Errorf("invalid recipient email: %w", err)
	}

	return nil
}
