package mailer

// Messenger describes the functions to deliver application specific messages.
type Messenger interface {
}

// Mailer implements the Messenger interface.
type Mailer struct {
	mailClient       MailClient
	confirmUrl       string
	resetPasswordUrl string
	from             string
	siteName         string
}

// ensure Mailer implements the Messenger interface.
var _ Messenger = &Mailer{}

// NewMailer creates a new mailer and returns the pointer.
func NewMailer(mailClient MailClient, mailFrom, siteName string) (*Mailer, error) {
	return &Mailer{
		mailClient: mailClient,
		from:       mailFrom,
		siteName:   siteName,
	}, nil
}
