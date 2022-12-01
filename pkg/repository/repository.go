package repository

import "form-to-mail-service/models"

type Delivery interface {
	SendEmail(customer *models.Customer) error
}

type Repository struct {
	Delivery
}

func NewRepository(config *SMTPConfig) *Repository {
	return &Repository{Delivery: NewDeliverySMTP(config)}
}
