package service

import (
	"fmt"
	"github.com/miguelpragier/handy"
	"processor/payment"
	"processor/payment/models"
	"sync"
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
	if req.Token == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("empty Token"))
		return false
	}

	var wg sync.WaitGroup
	cErr := make(chan *models.Error)

	go func() {
		var err models.Error
		wg.Add(1)
		defer wg.Done()
		s.validateRequest(req, &err)
		cErr <- &err
	}()

	var sData models.CardSensitiveData
	go func() {
		var err models.Error
		wg.Add(1)
		defer wg.Done()
		sData = s.repo.GetCardSensitiveData(req.Token, &err)
		cErr <- &err
	}()

	wg.Wait()
	req.Card.Sensitive = sData
	errors = <-cErr

	if len(errors.Validation) > 0 || len(errors.Internal) > 0 {
		return false
	}

	success := s.gate.ProcessPayment(req.Card, req.Process, req.AcquirerID, errors)

	if len(errors.Validation) > 0 || len(errors.Internal) > 0 {
		return false
	}

	return success
}

func (s *service) validateRequest(req models.Request, errors *models.Error) {

	if req.AcquirerID <= 0 {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid AcquirerID"))
		return
	}

	if req.Card.Open.Name == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("empty Card Name"))
		return
	}

	if req.Card.Open.Date == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("empty Card Date"))
		return
	}

	if req.Card.Open.Flag == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("empty Card Flag"))
		return
	}

	if req.Card.Sensitive.Number != "" || req.Card.Sensitive.CVV != "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid card info"))
		return
	}

	if req.Process.TotalValue <= 0 {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid Transaction Value"))
		return
	}

	var sum float64

	for _, item := range req.Process.Items {
		if item.Name == "" {
			errors.Validation = append(errors.Validation, fmt.Errorf("invalid Item Name"))
			return
		}

		if item.Value < 0 {
			errors.Validation = append(errors.Validation, fmt.Errorf("invalid Item Value"))
			return
		}

		sum += item.Value
	}

	if sum != req.Process.TotalValue {
		errors.Validation = append(errors.Validation, fmt.Errorf("the sum of items values is not equal to informed total value"))
		return
	}

	if req.Process.Installments <= 0 {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid number of installments"))
		return
	}

	if req.Process.Seller.Name == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid seller name"))
		return
	}

	if !handy.CheckCNPJ(req.Process.Seller.CNPJ) {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid seller CNPJ"))
		return
	}

	if req.Process.Seller.Address.Street == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid seller Adress"))
		return
	}

	if req.Process.Seller.Address.Number == 0 {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid seller Adress Number"))
		return
	}

	if req.Process.Seller.Address.ZipCode == "" {
		errors.Validation = append(errors.Validation, fmt.Errorf("invalid seller Adress Zip Code"))
		return
	}
}
