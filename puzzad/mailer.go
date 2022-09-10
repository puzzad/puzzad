package puzzad

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	htmlTemplate "html/template"
	"net/smtp"
	"path/filepath"
	textTemplate "text/template"

	"github.com/jordan-wright/email"
)

type MailClient struct {
	SMTPUser     string
	SMTPPassword string
	SMTPServer   string
	SMTPPort     int
	SMTPFrom     string
}

type EmailData struct {
	ValidateURL string
	Code        string
	Email       string
}

func (m *MailClient) SendPasswordResetLink(_ context.Context, address string, code string) error {
	data := EmailData{
		ValidateURL: "http://localhost:3000",
		Code:        code,
		Email:       address,
	}
	textData := bytes.NewBuffer(make([]byte, 0))
	htmlData := bytes.NewBuffer(make([]byte, 0))
	err := textTemplate.Must(textTemplate.New("plain.email").ParseFiles(filepath.Join("puzzad", "templates", "plain.email"))).Execute(textData, data)
	if err != nil {
		return err
	}
	err = htmlTemplate.Must(htmlTemplate.New("html.email").ParseFiles(filepath.Join("puzzad", "templates", "html.email"))).Execute(htmlData, data)
	if err != nil {
		return err
	}
	em := email.Email{
		ReplyTo: nil,
		From:    fmt.Sprintf("Puzzad Team <%s>", m.SMTPFrom),
		To:      []string{data.Email},
		Subject: "Puzzad Validation",
		Text:    textData.Bytes(),
		HTML:    htmlData.Bytes(),
		Sender:  fmt.Sprintf("Puzzad Team <%s>", m.SMTPFrom),
	}
	return em.SendWithStartTLS(fmt.Sprintf("%s:%d", m.SMTPServer, m.SMTPPort), smtp.PlainAuth("", m.SMTPUser, m.SMTPPassword, m.SMTPServer), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.SMTPServer,
	})
}
