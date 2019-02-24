package main

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

import gomail "gopkg.in/gomail.v2"

type MailJetMailBackend struct {
	d *gomail.Dialer
}

func NewMailJetMailBackend(config *MailConfig) *MailJetMailBackend {
	d := gomail.NewDialer(config.Server, config.Port, config.Email, config.Password)

	return &MailJetMailBackend{
		d: d,
	}
}

func (mail *MailJetMailBackend) SendMail(config *MailConfig, t *Template, params interface{}, subject string, to []string, attachments []string) error {
	data, err := ioutil.ReadAll(t.Data)
	if err != nil {
		return err
	}

	from := config.MetaData.FromName + " <" + config.MetaData.FromEmail + ">"
	body := new(bytes.Buffer)
	temp, err := template.New(t.Name).Parse(string(data))
	if err != nil {
		return err
	}
	err = temp.Execute(body, params)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	for i := range attachments {
		m.Attach(attachments[i])
	}

	if err := mail.d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
