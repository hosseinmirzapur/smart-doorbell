package services

import (
	"crypto/tls"
	"log"

	"github.com/hosseinmirzapur/sdb/config"
	"gopkg.in/gomail.v2"
)

type emailService struct{}

var dialer *gomail.Dialer

func NewEmailService() *emailService {
	email := config.GetEmailConfig()

	dialer = gomail.NewDialer(
		email.Host,
		email.Port,
		email.Username,
		email.Password,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &emailService{}

}

func (srv *emailService) Send(from, to, subject, message string) error {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	if err := dialer.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
