package mailsender

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type mailFile struct {
	MailHost   string   `json:"mailHost"`
	Port       string   `json:"port"`
	User       string   `json:"user"`
	Password   string   `json:"password"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
}

func TestMailSender_SecureSenderMail(t *testing.T) {
	//ports 587 465 80
	var fileName = "../mail.json"
	var mm mailFile
	file, err2 := ioutil.ReadFile(fileName)
	if err2 == nil {
		err := json.Unmarshal(file, &mm)
		fmt.Println("marshal err: ", err)
	}
	fmt.Println("file: ", mm)
	var ps SecureSender
	ps.User = mm.User
	ps.Password = mm.Password
	ps.MailHost = mm.MailHost
	ps.Port = mm.Port

	var m Mailer
	m.SenderAddress = mm.Sender
	m.Recipients = mm.Recipients
	m.Subject = "Test Mail"
	m.Body = "This is only a test."

	s := ps.GetNew()
	suc := s.SendMail(&m)
	fmt.Println("suc: ", suc)
	if !suc {
		t.Fail()
	}
}

func TestMailSender_SenderMail(t *testing.T) {
	//ports 587 465 80
	var fileName = "../mail.json"
	var mm mailFile
	file, err2 := ioutil.ReadFile(fileName)
	if err2 == nil {
		err := json.Unmarshal(file, &mm)
		fmt.Println("marshal err: ", err)
	}
	fmt.Println("file: ", mm)
	var ps PlainSender
	ps.User = mm.User
	ps.Password = mm.Password
	ps.MailHost = mm.MailHost
	ps.Port = mm.Port

	var m Mailer
	m.SenderAddress = mm.Sender
	m.Recipients = mm.Recipients
	m.Subject = "Test Mail"
	m.Body = "This is only a test."

	s := ps.GetNew()
	suc := s.SendMail(&m)
	fmt.Println("suc: ", suc)
	if suc {
		t.Fail()
	}
}

func TestMailSender_SenderMail2(t *testing.T) {
	//ports 587 465 80
	var fileName = "../mail3.json"
	var mm mailFile
	file, err2 := ioutil.ReadFile(fileName)
	if err2 == nil {
		err := json.Unmarshal(file, &mm)
		fmt.Println("marshal err: ", err)
	}
	fmt.Println("file: ", mm)
	var ps PlainSender
	ps.User = mm.User
	ps.Password = mm.Password
	ps.MailHost = mm.MailHost
	ps.Port = mm.Port

	var m Mailer
	m.SenderAddress = mm.Sender
	m.Recipients = mm.Recipients
	m.Subject = "Test Mail"
	m.Body = "This is only a test."

	s := ps.GetNew()
	suc := s.SendMail(&m)
	if !suc {
		t.Fail()
	}
}

func TestMailSender_SecureSenderMailOffice365(t *testing.T) {
	//ports 587 465 80
	var fileName = "../mail365.json"
	var mm mailFile
	file, err2 := ioutil.ReadFile(fileName)
	if err2 == nil {
		err := json.Unmarshal(file, &mm)
		fmt.Println("marshal err: ", err)
	}
	fmt.Println("file: ", mm)
	var ps Office365Sender
	ps.User = mm.User
	ps.Password = mm.Password
	ps.MailHost = mm.MailHost
	ps.Port = mm.Port

	var m Mailer
	m.SenderAddress = mm.Sender
	m.Recipients = mm.Recipients
	m.Subject = "Test Mail"
	m.Body = "This is only a test."

	s := ps.GetNew()
	suc := s.SendMail(&m)
	fmt.Println("suc: ", suc)
	if !suc {
		t.Fail()
	}
}
