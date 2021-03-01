package api

import (
	"bytes"
	"fmt"
	"net/smtp"

	embtemplate "github.com/magna5/embedded.template"
	log "github.com/magna5/go-logger"
)

//SendMail func to send device info
func SendMail(conf *Cfg, data Device) error {

	var auth smtp.Auth
	if conf.MailPass == "" || conf.SMTPPort == "25" {
		auth = smtp.PlainAuth("", conf.MailFrom, conf.MailPass, conf.SMTPHost)
	}

	body, err := loadTemplate(data)
	if err != nil {
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

func loadTemplate(data Device) (bytes.Buffer, error) {
	var body bytes.Buffer

	t, err := embtemplate.LoadTemplates()
	if err != nil {
		log.Error(err)
		return body, err
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: LM to connectwise sync details \n%s\n\n", mimeHeaders)))

	err = t.ExecuteTemplate(&body, "mail.tmpl", data)
	if err != nil {
		log.Error(err)
		return body, err
	}

	return body, nil
}
