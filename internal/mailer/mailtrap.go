package mailer

import (
	"bytes"
	"crypto/tls"
	"errors"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// 1) шаблон
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	var subjBuf, bodyBuf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&subjBuf, "subject", data); err != nil {
		return -1, err
	}
	if err := tmpl.ExecuteTemplate(&bodyBuf, "body", data); err != nil {
		return -1, err
	}

	// 2) письмо
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.fromEmail)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subjBuf.String())
	msg.AddAlternative("text/html", bodyBuf.String())

	// 3) SMTP-конфиг
	host := "live.smtp.mailtrap.io"
	port := 2525 // попробуй 2525 сначала
	user := "api"
	pass := m.apiKey // это ваш токен из Email Sending

	if isSandbox {
		host = "sandbox.smtp.mailtrap.io"
		port = 2525
		user = "96075a8743b851"
		pass = "2f792a505a51f5"
	}

	d := gomail.NewDialer(host, port, user, pass)
	d.TLSConfig = &tls.Config{ServerName: host}

	if err := d.DialAndSend(msg); err != nil {
		return -1, err
	}
	return 200, nil
}
