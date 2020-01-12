package gateway

import (
	"processor/core/requester"
	"processor/payment"
	"processor/payment/models"
)

type gateway struct {
	Repository payment.Repository
	Requester  requester.Requester
}

func NewGateway(r payment.Repository, req requester.Requester) payment.Gateway {
	return &gateway{
		Repository: r,
		Requester:  req,
	}
}

func (g *gateway) ProcessPayment(card models.CardData, process models.Process, acquirerID int64, errors *models.Error) bool {
	url := g.Repository.GetAcquirerURL(acquirerID, errors)

	if len(errors.Internal) > 0 || len(errors.Validation) > 0 {
		return false
	}

	success := g.Requester.SendBuyRequest(process, card, url, errors)

	if len(errors.Internal) > 0 || len(errors.Validation) > 0 {
		return false
	}

	return success
}
