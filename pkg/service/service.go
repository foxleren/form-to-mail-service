package service

import (
	"form-to-mail-service/models"
	"form-to-mail-service/pkg/repository"
)

type Delivery interface {
	SendEmail(customer *models.Customer) error
}

type Service struct {
	Delivery
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Delivery: NewDeliveryService(repos.Delivery),
	}
}
