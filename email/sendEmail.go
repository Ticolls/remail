package email

func SendEmail(subject string, message string) error {

	// Set E-Mail subject
	mail.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	mail.SetBody("text/plain", message)

	// Now send E-Mail
	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
