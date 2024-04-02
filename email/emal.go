package email

import (
	"log"

	"gopkg.in/mail.v2"
)

/*
Метод отправляет email
*/
func Send(dialSettings *mail.Dialer, from string, to []string, cc []string, subject, body string, attachmetns ...string) {

	mail := mail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	mail.SetHeader("Cc", cc...)

	if len(attachmetns) != 0 {
		for _, file := range attachmetns {

			mail.Attach(file)
		}
	}

	if err := dialSettings.DialAndSend(mail); err != nil {
		log.Println(err)
	}

}
