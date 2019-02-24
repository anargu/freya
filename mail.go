package main

func (f *Freya) SendMail(templateName string, params interface{}, subject string, to []string, attachments []string) error {

	template, err := f.GetTemplateByName(templateName, true)
	if err != nil {
		return err
	}

	err = f.MailBackend.SendMail(f.Config.Mail, template, params, subject, to, attachments)
	if err != nil {
		return err
	}

	return nil
}
