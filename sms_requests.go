package main

type PhoneNumber struct {
	CountryCode string
	Number      string
}

type SMSSendResponse string

type SMSBackend interface {
	SendSMS(config *SMSConfig, to *PhoneNumber, t *Template, params map[string]string) (SMSSendResponse, error)
	GetBalance() (float64, error)
}
