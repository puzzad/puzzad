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
	URL          string
}

type EmailData struct {
	URL   string
	Code  string
	Email string
}

func (m *MailClient) SendEmailVerifyLink(_ context.Context, address string, code string) error {
	data := EmailData{
		URL:   m.URL,
		Code:  code,
		Email: address,
	}
	return m.sendEmail("verify", data)
}

func (m *MailClient) SendPasswordResetLink(_ context.Context, address string, code string) error {
	data := EmailData{
		URL:   m.URL,
		Code:  code,
		Email: address,
	}
	return m.sendEmail("passreset", data)
}

func (m *MailClient) sendEmail(template string, data EmailData) error {
	textData := bytes.NewBuffer(make([]byte, 0))
	htmlData := bytes.NewBuffer(make([]byte, 0))
	err := textTemplate.Must(textTemplate.New(template+".txt").ParseFiles(filepath.Join("puzzad", "templates", template+".txt"))).Execute(textData, data)
	if err != nil {
		return err
	}
	err = htmlTemplate.Must(htmlTemplate.New(template+".html").ParseFiles(filepath.Join("puzzad", "templates", template+".html"))).Execute(htmlData, data)
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
