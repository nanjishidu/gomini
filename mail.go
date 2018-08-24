package gomini

import (
	"crypto/tls"
	"log"
	"time"

	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	host        string
	username    string
	password    string
	port        int
	dialer      *gomail.Dialer
	messageChan chan *gomail.Message
}

func NewMail(host string, port int, username, password string, num ...int64) *Mail {
	var messageChanBuffer int64
	if len(num) > 0 {
		messageChanBuffer = num[0]
	} else {
		messageChanBuffer = 1
	}
	return &Mail{
		host:        host,
		username:    username,
		password:    password,
		port:        port,
		dialer:      gomail.NewDialer(host, port, username, password),
		messageChan: make(chan *gomail.Message, messageChanBuffer),
	}
}
func (m *Mail) SetMessage(params map[string]interface{}) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.username)
	if v, ok := params["to"]; ok {
		msg.SetHeader("To", v.([]string)...)
	}
	if v, ok := params["cc"]; ok {
		msg.SetHeader("Cc", v.([]string)...)
	}
	if v, ok := params["subject"]; ok {
		msg.SetHeader("Subject", v.(string))
	}
	if v, ok := params["text"]; ok {
		msg.SetBody("text/html", v.(string))
	}
	if v, ok := params["html"]; ok {
		msg.SetBody("text/html", v.(string))
	}
	if v, ok := params["attachFile"]; ok {
		for _, v2 := range v.([]string) {
			msg.Attach(v2)
		}
	}
	m.messageChan <- msg
}
func (m *Mail) SendEmailDaemon() {
	var (
		s    gomail.SendCloser
		err  error
		open = false
	)
	for {
		select {
		case msg, ok := <-m.messageChan:
			if !ok || msg == nil {
				return
			}
			if !open {
				if s, err = m.dialer.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			if err := gomail.Send(s, msg); err != nil {
				log.Print(err)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(30 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}
func (m *Mail) SendEmailOnce() error {
	m.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if msg, ok := <-m.messageChan; ok && msg != nil {
		return m.dialer.DialAndSend(msg)
	}
	return nil
}
