package mailer

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Messenger describes the functions to deliver application specific messages.
type Messenger interface {
	Contact(name, subject, mail, content, lang string) error
}

// Mailer implements the Messenger interface.
type Mailer struct {
	mailClient MailClient
	contactTo  string
	from       string
	siteName   string
}

// ensure Mailer implements the Messenger interface.
var _ Messenger = &Mailer{}

// NewMailer creates a new mailer and returns the pointer.
func NewMailer(mailClient MailClient, mailFrom, siteName, contactTo string) (*Mailer, error) {
	return &Mailer{
		mailClient: mailClient,
		from:       mailFrom,
		siteName:   siteName,
		contactTo:  contactTo,
	}, nil
}

// Contact sends the message to the visitor and the site owner.
func (m Mailer) Contact(name, subject, mail, content, lang string) error {
	printer := message.NewPrinter(language.Make(lang))
	err := m.mailClient.Send(
		m.from,
		[]string{m.contactTo},
		fmt.Sprintf("%s: %s - %s", m.siteName, name, subject),
		fmt.Sprintf("%s <%s>:\n%s", name, mail, content),
		nil,
		nil,
	)
	if err != nil {
		status.Errorf(codes.Internal, printer.Sprintf("error while sending mail: %s", err))
	}
	return nil
}
