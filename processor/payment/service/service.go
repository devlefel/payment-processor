package service

import (
	"processor/payment"
	"processor/payment/models"
)

type service struct {
	repo payment.Repository
	gate payment.Gateway
}

func NewService(r payment.Repository, g payment.Gateway) payment.Service {
	return &service{
		repo: r,
		gate: g,
	}
}

func (s *service) ProcessPayment(req models.Request, errors *models.Error) bool {
	return true
}
