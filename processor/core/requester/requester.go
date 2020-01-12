package requester

import "processor/payment/models"

type Requester interface {
	SendBuyRequest(process models.Process, card models.CardData, acquirerURL string, errors *models.Error) bool
}
