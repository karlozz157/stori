package mailers

type Mailer interface {
	SendEmail(emailData EmailData) error
}

type EmailData struct {
	To       []string
	Subject  string
	Template string
	Data     interface{}
}
