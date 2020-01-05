package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"processor/payment/mocks"
	"processor/payment/models"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repo := mocks.NewMockRepository(controller)
	gate := mocks.NewMockGateway(controller)

	service := NewService(repo, gate)

	test := struct {
		Input          models.Request
		ExpectedResult bool
		ExpectedError  models.Error
	}{
		Input: models.Request{
			Token: "abc1235efrda54fea3f1ea6f8rafeghr564",
			Card: models.CardData{
				Open: models.CardOpenData{
					Name: "Felipe Gomes",
					Flag: "VISA",
					Date: "09/2023",
				},
				Sensitive: models.CardSensitiveData{
					Number: "",
					CVV:    "",
				},
			},
			Process: models.Process{
				TotalValue: 1000.00,
				Items: []models.Item{
					{Name: "Geladeira", Value: 1000.00},
				},
				Installments: 12,
				Seller: models.Seller{
					Name: "Magazine Luiza",
					CNPJ: "15.958.655/0001-44",
					Address: models.SellerAddress{
						Street:  "Rua Jatobá",
						Number:  255,
						ZipCode: "04812190",
					},
				},
			},
			AcquirerID: 2,
		},
		ExpectedResult: true,
		ExpectedError: models.Error{
			Validation: nil,
			Internal:   nil,
		},
	}

	var err models.Error
	repo.EXPECT().GetCardSensitiveData(test.Input.Token, &err).Return(models.CardSensitiveData{Number: "1234567891234567", CVV: ""})
	gate.EXPECT().ProcessPayment(test.Input.Card, test.Input.Process, test.Input.AcquirerID, err).Return(test.ExpectedResult)

	ret := service.ProcessPayment(test.Input, &err)

	if assert.Equal(t, test.ExpectedError, err, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, err) {
		assert.Equal(t, test.ExpectedResult, ret, "We expected return: %#v but got: %#v", test.ExpectedResult, ret)
	}
}
