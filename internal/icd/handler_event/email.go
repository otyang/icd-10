package handler_event

import (
	"crypto/tls"
	"fmt"

	"github.com/gookit/event"
	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/logger"

	gomail "gopkg.in/mail.v2"
)

type Handler struct {
	Log    logger.Interface
	Config *config.Config
}

func NewHandler(config *config.Config, Log logger.Interface) *Handler {
	return &Handler{
		Log:    Log,
		Config: config,
	}
}

var m = gomail.NewMessage() // lets do it here so we can cache the initialisation

func (h *Handler) EventHandlerFileUpload(e event.Event) error {

	toEmail := e.Data()["0"].(string)

	m.SetHeader("To", toEmail)
	m.SetHeader("From", h.Config.SMTP.Email)
	m.SetHeader("Subject", "ICD CSV File Upload Completed")
	m.SetBody("text/plain", "Hello\n\n Your file upload was completed succesfully.\n\n Admin")

	// Settings for SMTP server
	d := gomail.NewDialer(h.Config.SMTP.Server, h.Config.SMTP.Port, h.Config.SMTP.Email, h.Config.SMTP.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: h.Config.SMTP.EnableTLS} // In production  set to false. In development set to true

	// Now send E-Mail
	// if err := d.DialAndSend(m); err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(toEmail)

	return nil
}
