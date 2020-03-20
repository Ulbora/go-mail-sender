package mailsender

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

//Mailer Mailer
type Mailer struct {
	SenderAddress string
	Recipients    []string
	Subject       string
	Body          string
}

//Sender Sender
type Sender interface {
	SendMail(mailer *Mailer) bool
}

//PlainSender PlainSender
type PlainSender struct {
	User     string
	Password string
	MailHost string
	Port     string
}

//SendMail SendMail
func (m *PlainSender) SendMail(mailer *Mailer) bool {
	var rtn bool
	if mailer.Recipients != nil && len(mailer.Recipients) > 0 {
		auth := smtp.PlainAuth("", m.User, m.Password, m.MailHost)
		var recips = "To: "
		for r := range mailer.Recipients {
			recips += mailer.Recipients[r]
			if r < len(mailer.Recipients)-1 {
				recips += ";"
			}
		}
		recips += "\r\n"
		var msg = []byte(
			"From: " + mailer.SenderAddress + "\r\n" +
				recips +
				"Subject: " + mailer.Subject + "\r\n" +
				"\r\n" +
				mailer.Body + "\r\n")
		//fmt.Println("mailer: ", *mailer)
		err := smtp.SendMail(m.MailHost+":"+m.Port, auth, mailer.SenderAddress, mailer.Recipients, msg)
		if err != nil {
			log.Println("Mail send error: ", err)
		} else {
			rtn = true
		}
	}
	return rtn
}

//GetNew GetNew
func (m *PlainSender) GetNew() Sender {
	var s Sender
	s = m
	return s
}

//SecureSender SecureSender
type SecureSender struct {
	User     string
	Password string
	MailHost string
	Port     string
}

//SendMail SendMail
func (m *SecureSender) SendMail(mailer *Mailer) bool {
	var rtn bool
	if mailer.Recipients != nil && len(mailer.Recipients) > 0 {
		var hostPort = m.MailHost + ":" + m.Port
		var tlsCfg tls.Config
		tlsCfg.InsecureSkipVerify = true
		tlsCfg.ServerName = m.MailHost
		auth := smtp.PlainAuth("", m.User, m.Password, m.MailHost)
		conn, derr := tls.Dial("tcp", hostPort, &tlsCfg)
		if derr == nil {
			client, cerr := smtp.NewClient(conn, m.MailHost)
			if cerr == nil {
				aerr := client.Auth(auth)
				if aerr == nil {
					merr := client.Mail(m.User)
					if merr == nil {
						var recips = "To: "
						for r := range mailer.Recipients {
							recips += mailer.Recipients[r]
							if r < len(mailer.Recipients)-1 {
								recips += ";"
							}
							rerr := client.Rcpt(mailer.Recipients[r])
							if rerr != nil {
								log.Println("client Rcpt Error: ", rerr)
							}
						}
						recips += "\r\n"
						//fmt.Println("recips", recips)
						w, derr := client.Data()
						if derr == nil {
							var msg = []byte(
								"From: " + mailer.SenderAddress + "\r\n" +
									recips +
									"Subject: " + mailer.Subject + "\r\n" +
									"\r\n" +
									mailer.Body + "\r\n")
							_, werr := w.Write(msg)
							if werr == nil {
								cerr := w.Close()
								if cerr == nil {
									rtn = true
								} else {
									log.Println("client Close Error: ", cerr)
								}
							} else {
								log.Println("client Write Error: ", werr)
							}
						} else {
							log.Println("client Data Error: ", derr)
						}
					} else {
						log.Println("client Mail Error: ", merr)
					}
				} else {
					log.Println("client Auth Error: ", aerr)
				}
				client.Quit()
			} else {
				log.Println("client Error: ", cerr)
			}
		} else {
			log.Println("Dial Error: ", derr)
		}
	}
	return rtn
}

//GetNew GetNew
func (m *SecureSender) GetNew() Sender {
	var s Sender
	s = m
	return s
}

//go mod init github.com/Ulbora/go-mail-sender
