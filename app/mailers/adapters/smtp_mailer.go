package mailers_adapaters

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/karlozz157/stori/app/mailers"
)

type StmpMailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService() mailers.Mailer {
	return &StmpMailer{}
}

func (s *StmpMailer) SendEmail(emailData mailers.EmailData) error {
	template, err := s.getTemplate(emailData)
	if err != nil {
		return err
	}

	message := []byte(
		fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
			emailData.To,
			emailData.Subject,
			template))

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", s.Host, s.Port), auth, s.Username, emailData.To, message)

	if err != nil {
		log.Printf("error sending email: %s", err)
	}

	return err
}

func (e *StmpMailer) getTemplate(emailData mailers.EmailData) (string, error) {
	file, err := os.Open(emailData.Template)
	if err != nil {
		return "", fmt.Errorf("error opening template file: %s", err)
	}
	defer file.Close()

	template, err := template.ParseFiles(emailData.Template)
	if err != nil {
		return "", fmt.Errorf("error parsing template file: %s", err)
	}

	var body bytes.Buffer
	err = template.Execute(&body, emailData.Data)
	if err != nil {
		return "", fmt.Errorf("error rendering template: %s", err)
	}

	return body.String(), nil
}
