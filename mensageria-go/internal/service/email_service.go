package service

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"

	"mensageria-go/internal/config"
	"mensageria-go/internal/model"
)

type EmailService struct {
	cfg config.Config
}

func NewEmailService(cfg config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func (e *EmailService) EnviarEmail(
	destino string,
	data model.EmailData,
) error {

	tmpl, err := template.ParseFiles("templates/pedido_confirmado.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		e.cfg.EmailUser,
		e.cfg.EmailPass,
		e.cfg.EmailHost,
	)

	msg := ""
	msg += "MIME-version: 1.0;\n"
	msg += "Content-Type: text/html; charset=\"UTF-8\";\n"
	msg += "Subject: ✅ Seu Pedido foi Confirmado - Neoshop\n\n"
	msg += body.String()

	err = smtp.SendMail(
		e.cfg.EmailHost+":"+e.cfg.EmailPort,
		auth,
		e.cfg.EmailUser,
		[]string{destino},
		[]byte(msg),
	)

	if err != nil {
		return err
	}

	log.Println("📧 Email enviado para:", destino)
	return nil
}