package contracts

// Notifier defines a minimal contract for sending messages to recipients.
type Notifier interface {
	Send(to, subject, message string) error
}
