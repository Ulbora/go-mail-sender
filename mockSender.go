package mailsender

//MockPlainSender MockPlainSender
type MockPlainSender struct {
	User        string
	Password    string
	MailHost    string
	Port        string
	MockSuccess bool
}

//SendMail SendMail
func (m *MockPlainSender) SendMail(mailer *Mailer) bool {
	return m.MockSuccess
}

//GetNew GetNew
func (m *MockPlainSender) GetNew() Sender {
	var s Sender
	s = m
	return s
}

//MockSecureSender SecureSender
type MockSecureSender struct {
	User        string
	Password    string
	MailHost    string
	Port        string
	MockSuccess bool
}

//SendMail SendMail
func (m *MockSecureSender) SendMail(mailer *Mailer) bool {
	return m.MockSuccess
}

//GetNew GetNew
func (m *MockSecureSender) GetNew() Sender {
	var s Sender
	s = m
	return s
}
