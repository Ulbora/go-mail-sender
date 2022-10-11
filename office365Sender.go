package mailsender

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/smtp"
)

//auth365 auth365
type auth365 struct {
	username string
	password string
}

//Auth365 Auth365
func Auth365(username, password string) smtp.Auth {
	return &auth365{username, password}
}

//Start Start
func (a *auth365) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

//Next Next
func (a *auth365) Next(fromServer []byte, more bool) ([]byte, error) {
	var rtn []byte
	var rtnErr error
	if more {
		switch string(fromServer) {
		case "Username:":
			rtn = []byte(a.username)
		case "Password:":
			rtn = []byte(a.password)
		default:
			rtn = nil
			rtnErr = errors.New("Unknown from server")
		}
	}
	return rtn, rtnErr
}

//Office365Sender Office365Sender
type Office365Sender struct {
	User     string
	Password string
	MailHost string
	Port     string
}

//SendMail SendMail
func (m *Office365Sender) SendMail(mailer *Mailer) bool {
	var rtn bool
	if mailer.Recipients != nil && len(mailer.Recipients) > 0 {
		var hostPort = m.MailHost + ":" + m.Port
		conn, oerr := net.Dial("tcp", hostPort)
		if oerr == nil {
			client, cerr := smtp.NewClient(conn, m.MailHost)
			if cerr == nil {
				//log.Println(client)
				var tlsCfg tls.Config
				tlsCfg.ServerName = m.MailHost
				sterr := client.StartTLS(&tlsCfg)
				if sterr == nil {
					//auth := smtp.Auth()
					auth := Auth365(m.User, m.Password) //(m.User, m.Password)
					aerr := client.Auth(auth)
					if aerr == nil {
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
						if err == nil {
							rtn = true
						} else {
							log.Println("Mail send error: ", err)
						}
					} else {
						log.Println("client Auth Error: ", aerr)
					}
				} else {
					log.Println("client StartTLS Error: ", sterr)
				}
			} else {
				log.Println("client Error: ", cerr)
			}
		} else {
			log.Println("Dial Error: ", oerr)
		}
	}

	return rtn
}

//GetNew GetNew
func (m *Office365Sender) GetNew() Sender {
	var s Sender
	s = m
	return s
}
