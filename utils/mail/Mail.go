package mail

import (
	"gopkg.in/gomail.v2"
)

const (
	username string = "admin@gaiamount.com"
	password string = "Shengdan1225"
	host     string = "smtp.mxhichina.com"
	fromName string = "Gaiamount"
	port     int    = 465
)

func Send(to, html, subject string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", username, fromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)
	d := gomail.NewDialer(host, port, username, password)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
