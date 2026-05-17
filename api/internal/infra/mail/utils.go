package mail

import (
	"fmt"
	"html"
)

func plainMessageBody(code string) string {
	return fmt.Sprintf("Код авторизации: %s. Код действителен в течение 5 минут.", code)
}

func htmlMessageBody(code string) string {
	return fmt.Sprintf(
		"Код авторизации: <b>%s</b>. Код действителен в течение 5 минут.",
		html.EscapeString(code),
	)
}
