package payment

import "processor/payment/models"

type Gateway interface {
	ProcessPayment(card models.CardData, process models.Process, acquirerID int64, errors *models.Error) bool
}
