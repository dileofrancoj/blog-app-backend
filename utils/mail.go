package utils

import (
	"crypto/tls"

	"os"

	gomail "gopkg.in/gomail.v2"
)

type smtpServer struct {
 host string
 port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
 return s.host + ":" + s.port
}

func SendMail(to string, subject string, body string) (bool, error){
	
	mail := gomail.NewDialer(os.Getenv("MAIL_SERVER"), 587, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))
	mail.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := mail.DialAndSend(m); err != nil {
		panic(err)
	}

	return true, nil
}
