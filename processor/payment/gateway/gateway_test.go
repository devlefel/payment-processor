package gateway

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	coreMocks "processor/core/mocks"
	"processor/payment/mocks"
	"processor/payment/models"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repo := mocks.NewMockRepository(controller)
	requester := coreMocks.NewMockRequester(controller)
	gate := NewGateway(repo, requester)

	type test struct {
		InputCard      models.CardData
		InputProcess   models.Process
		AcquirerID     int64
		ExpectedURL    string
		ExpectedReturn bool
		ExpectedError  models.Error
	}

	var tests = make(map[string]test)

	tests["ReturnValue"] = test{
		ExpectedReturn: true,
		ExpectedURL:    "http://api.stone.com.br/process",
		InputCard: models.CardData{
			Open: models.CardOpenData{
				Name: "Felipe Gomes",
				Flag: "VISA",
				Date: "03/2023",
			},
			Sensitive: models.CardSensitiveData{
				Number: "123456789123456",
				CVV:    "123",
			},
		},
		AcquirerID: 1,
		ExpectedError: models.Error{
			Validation: nil,
			Internal:   nil,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			var errors = test.ExpectedError
			repo.EXPECT().GetAcquirerURL(test.AcquirerID, &errors).Return(test.ExpectedURL)
			requester.EXPECT().SendBuyRequest(test.InputProcess, test.InputCard, test.ExpectedURL, &errors).Return(test.ExpectedReturn)
			ret := gate.ProcessPayment(test.InputCard, test.InputProcess, test.AcquirerID, &errors)

			if assert.Equal(t, test.ExpectedError, errors, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, errors) {
				assert.Equal(t, test.ExpectedReturn, ret, "We expected return: %#v but got: %#v", test.ExpectedReturn, ret)
			}
		})
	}
}
