package repository

import (
	"fmt"
	"form-to-mail-service/models"
	"net/smtp"
)

type DeliverySMTP struct {
	config *SMTPConfig
}

func NewDeliverySMTP(config *SMTPConfig) *DeliverySMTP {
	return &DeliverySMTP{config: config}
}

func (p *DeliverySMTP) SendEmail(customer *models.Customer) error {
	from, to := p.config.EmailAddress, p.config.EmailAddress
	password := p.config.Password

	auth := smtp.PlainAuth("", from, password, p.config.Host)

	email := buildEmail(to, customer)

	if err := smtp.SendMail(p.config.Host+":"+p.config.Port, auth, from, []string{to}, []byte(email)); err != nil {
		return err
	}
	
	return nil
}

func buildEmail(to string, customer *models.Customer) string {
	name := customer.Name
	if len(name) == 0 {
		name = "Unknown customer"
	}

	return fmt.Sprintf("To: %s\r\n"+
		"Subject: Request from customer\r\n"+
		"\r\n"+
		"Name: %s\r\n"+
		"Phone: %s\r\n", to, name, customer.Phone)
}
