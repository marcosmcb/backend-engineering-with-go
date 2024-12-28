package mailer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"time"

	"github.com/resend/resend-go/v2"
)

type ResendMailer struct {
	fromEmail string
	apiKey    string
	client    *resend.Client
}

func NewResend(apiKey, fromEMail string) *ResendMailer {
	client := resend.NewClient(apiKey)
	return &ResendMailer{
		fromEmail: fromEMail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (m *ResendMailer) Send(templateFile, username, email string, data any, isSandbox bool) (string, error) {
	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return "", err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return "", err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return "", err
	}

	params := &resend.SendEmailRequest{
		To:      []string{email},
		From:    m.fromEmail,
		Html:    body.String(),
		Subject: subject.String(),
	}

	var retryErr error
	for i := 0; i < maxRetries; i++ {
		response, retryErr := m.client.Emails.SendWithContext(context.TODO(), params)
		if retryErr != nil {
			// Exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		return response.Id, nil
	}
	return "", fmt.Errorf("failed to send email after %d attempts, error %v", maxRetries, retryErr)
}
