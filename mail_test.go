package gomini

import (
	"testing"
	"time"
)

var (
	host     = "smtp.exmail.qq.com"
	port     = 465
	username = ""
	password = ""
	to       = ""
	params   = map[string]interface{}{
		"to":         []string{to},
		"subject":    "邮件发送",
		"html":       "您正在进行邮件发送测试,请忽略。。。",
		"attachFile": []string{"/tmp/201808/5b5fadbc13389e93cf72502dec1208b9.png"},
	}
	m = NewMail(host, port, username, password)
)

func TestSendMail(t *testing.T) {
	m.SetMessage(params)
	if err := m.SendEmailOnce(); err != nil {
		t.Error(err)
	}
}
func TestSendMailDaemon(t *testing.T) {
	go m.SendEmailDaemon()
	for i := 0; i < 5; i++ {
		m.SetMessage(params)
	}
	time.Sleep(3 * time.Second)
}
