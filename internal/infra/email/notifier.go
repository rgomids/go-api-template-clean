package email

import "log"

// EmailNotifier defines an interface for sending emails.
type EmailNotifier interface {
	Send(to, subject, body string) error
}

// SMTPEmailNotifier is a simple notifier that logs email sending.
type SMTPEmailNotifier struct{}

// NewSMTPEmailNotifier creates a new SMTPEmailNotifier instance.
func NewSMTPEmailNotifier() *SMTPEmailNotifier {
	return &SMTPEmailNotifier{}
}

// Send logs the email sending operation.
func (n *SMTPEmailNotifier) Send(to, subject, body string) error {
	log.Printf("sending email to %s with subject %s: %s", to, subject, body)
	return nil
}

var _ EmailNotifier = (*SMTPEmailNotifier)(nil)
