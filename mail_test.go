package gomini

var (
	host     = "smtp.exmail.qq.com"
	port     = 465
	username = ""
	password = ""
	to       = ""
	params   = map[string]interface{}{
		"to":          []string{to},
		"subject":     "邮件发送",
		"html":        "您正在进行邮件发送测试,请忽略。。。",
		"attach":      []string{"/tmp/201808/5b5fadbc13389e93cf72502dec1208b9.png"},
		"embed":       []string{"/tmp/201808/5b5fadbc13389e93cf72502dec1208b9.png"},
		"alternative": "<div style='margin-top:20px'><br>-----------------------<br><b>nanjishidu</b></div>",
	}
	m = NewMail(host, port, username, password)
)

//func TestSendMail(t *testing.T) {
//	m.SetMessage(params)
//	if err := m.SendEmailOnce(); err != nil {
//		t.Error(err)
//	}
//}
//func TestSendMailDaemon(t *testing.T) {
//	go m.SendEmailDaemon()
//	for i := 0; i < 1; i++ {
//		m.SetMessage(params)
//	}
//	time.Sleep(3 * time.Second)
//}
