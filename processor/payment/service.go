package payment

import "processor/payment/models"

type Service interface {
	ProcessPayment(req models.Request, errors *models.Error) bool
}
