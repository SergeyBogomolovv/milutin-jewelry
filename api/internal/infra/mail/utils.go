package mail

import "fmt"

func messageBody(code string) string {
	return fmt.Sprintf("Код авторизации: <b>%s</b>. Код действителен в течении 5 минут", code)
}
