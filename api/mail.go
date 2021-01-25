package api

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	log "github.com/magna5/go-logger"
)

//SendMail func to send device info
func SendMail(conf *Cfg, data Device) error {

	auth := smtp.PlainAuth("", conf.MailFrom, conf.MailPass, conf.SMTPHost)
	t, err := template.ParseFiles("api/template.html")
	if err != nil {
		log.Error(err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: LM to connectwise sync details \n%s\n\n", mimeHeaders)))

	err = t.Execute(&body, data)
	if err != nil {
		log.Error(err)
		return err
	}

	err = smtp.SendMail(conf.SMTPHost+":"+conf.SMTPPort, auth, conf.MailFrom, conf.MailTo, body.Bytes())
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Email Sent!")

	return err

}
