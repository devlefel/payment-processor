package requester

import (
	"fmt"
	"processor/payment/models"
	"strings"
)

//this is the struct that will implement the requester, but since is a mock we dont have nothing on this struct
type requester struct{}

// Returns a new Requester that in theory would call external acquirers API to send a process, but since these API are a mock on this challenge we just pretend that we call the requests but instead we return some mocked responses
func NewRequester() Requester {
	return &requester{}
}

//SendBuyRequest Implementation here is a Mocked version, because from here on its far from the challenge's scope
func (r *requester) SendBuyRequest(process models.Process, card models.CardData, acquirerURL string, errors *models.Error) bool {
	if process.TotalValue > 1000.00 {
		errors.Validation = append(errors.Validation, fmt.Errorf("total value is above allowed"))
		return false
	}

	if strings.Contains(card.Open.Name, "João Antônio") {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid credit card data"))
		return false
	}

	if strings.Contains(acquirerURL, "cielo") {
		errors.Validation = append(errors.Internal, fmt.Errorf("error sending process to aquire"))
		return false
	}

	return true
}
