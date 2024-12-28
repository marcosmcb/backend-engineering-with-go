package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

func NewSendgrid(apiKey, fromEMail string) *SendgridMailer {
	client := sendgrid.NewSendClient(apiKey)
	return &SendgridMailer{
		fromEmail: fromEMail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (m *SendgridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}
	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())

	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox,
		},
	})

	for i := 0; i < maxRetries; i++ {
		response, err := m.client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %v, attempt %d of %d\n", email, i+1, maxRetries)
			log.Printf("Error: %v\n", err.Error())

			// Exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		log.Printf("Email sent with status code %v\n", response.StatusCode)
		return nil
	}
	return fmt.Errorf("failed to send email after %d attempts", maxRetries)
}