package mailsender

//Mock365Sender Mock365Sender
type Mock365Sender struct {
	User        string
	Password    string
	MailHost    string
	Port        string
	MockSuccess bool
}

//SendMail SendMail
func (m *Mock365Sender) SendMail(mailer *Mailer) bool {
	return m.MockSuccess
}

//GetNew GetNew
func (m *Mock365Sender) GetNew() Sender {
	var s Sender
	s = m
	return s
}
