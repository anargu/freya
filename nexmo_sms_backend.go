package main

import (
	"bytes"
	"gopkg.in/njern/gonexmo.v2"
	"io/ioutil"
	"text/template"
)

type NexmoSMSBackend struct {
	client *nexmo.Client
}

func NewNexmoSMSBackend(config *SMSConfig) (*NexmoSMSBackend, error) {
	nexmoClient, err := nexmo.NewClient(config.Key, config.Secret)
	if err != nil {
		return nil, err
	}

	return &NexmoSMSBackend{
		client: nexmoClient,
	}, nil
}

func (n *NexmoSMSBackend) SendSMS(config *SMSConfig, to *PhoneNumber, t *Template, params interface{}) (SMSSendResponse, error) {
	data, err := ioutil.ReadAll(t.Data)
	if err != nil {
		return "", err
	}

	temp, err := template.New("nexmo").Parse(string(data))

	buffer := new(bytes.Buffer)

	if err = temp.Execute(buffer, data); err != nil {
		return "", err
	}

	message := &nexmo.SMSMessage{
		From:  config.From,
		To:    to.CountryCode + to.Number,
		Body:  buffer.Bytes(),
		Type:  nexmo.Text,
		Class: nexmo.Standard,
	}

	_, err = n.client.SMS.Send(message)
	if err != nil {
		return "", err
	}

	return "", nil
}
