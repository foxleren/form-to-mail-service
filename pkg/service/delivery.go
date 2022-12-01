package service

import (
	"form-to-mail-service/models"
	"form-to-mail-service/pkg/repository"
)

type DeliveryService struct {
	repo repository.Delivery
}

func NewDeliveryService(repo repository.Delivery) *DeliveryService {
	return &DeliveryService{repo: repo}
}

func (s DeliveryService) SendEmail(customer *models.Customer) error {
	return s.repo.SendEmail(customer)
}
