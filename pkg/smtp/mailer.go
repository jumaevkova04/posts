package smtp

import (
	"time"

	"github.com/go-mail/mail/v2"
)

type Mailer struct {
	dialer *mail.Dialer
	from   string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 20 * time.Second

	return &Mailer{
		dialer: dialer,
		from:   from,
	}
}

// Send - send an email notification
func (m *Mailer) Send(receiver string, subject, payload string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", payload)

	return m.dialer.DialAndSend(msg)
}

// SendBulk - send notification to multiple emails
func (m *Mailer) SendBulk(receivers []string, subject, payload string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", receivers...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", payload)

	return m.dialer.DialAndSend(msg)
}

// SendWithAttachment - sends notification via email attachment should be a path to a file
func (m *Mailer) SendWithAttachment(receiver, subject, payload, attachment string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", payload)
	msg.Attach(attachment)

	return m.dialer.DialAndSend(msg)
}
