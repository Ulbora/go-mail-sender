package mailsender

import (
	"testing"
)

func TestMockSender_SendPlain(t *testing.T) {
	var ms MockPlainSender
	ms.MockSuccess = true
	s := ms.GetNew()
	var m Mailer
	m.SenderAddress = "sender@gmail.com"
	m.Recipients = []string{"you@gmail.com"}
	m.Subject = "Test Mail"
	m.Body = "This is only a test."
	suc := s.SendMail(&m)
	if !suc {
		t.Fail()
	}
}

func TestMockSender_SendSecure(t *testing.T) {
	var ms MockSecureSender
	ms.MockSuccess = true
	s := ms.GetNew()
	var m Mailer
	m.SenderAddress = "sender@gmail.com"
	m.Recipients = []string{"you@gmail.com"}
	m.Subject = "Test Mail"
	m.Body = "This is only a test."
	suc := s.SendMail(&m)
	if !suc {
		t.Fail()
	}
}
