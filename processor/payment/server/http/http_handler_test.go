package http

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"processor/payment/mocks"
	"processor/payment/models"
	"testing"
)

type test struct {
	Input          models.Request
	StatusCode     int
	ExpectedOutput bool
	ExpectedError  models.Error
}

func TestProcessPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	r := chi.NewRouter()
	NewHandler(service, r)

	w := httptest.NewRecorder()

	theTest := test{
		StatusCode:     http.StatusOK,
		ExpectedOutput: true,
		ExpectedError: models.Error{
			Validation: nil,
			Internal:   nil,
		},
		Input: models.Request{
			Token: "56r4ea21a64fae1fa3efaef12ae35f",
			Card: models.CardData{
				Open: models.CardOpenData{
					Name: "Felipe Gomes",
					Flag: "VISA",
					Date: "09/23",
				},
				Sensitive: models.CardSensitiveData{
					Number: "1234567891234567",
					CVV:    "123",
				},
			},
			Process: models.Process{
				TotalValue: 1535.25,
				Items: []models.Item{
					{
						Name:  "Geladeira",
						Value: 1535.25,
					},
				},
				Installments: 12,
				Seller: models.Seller{
					Name: "Magazine Luiza",
					CNPJ: "90.095.874/0001-62",
					Address: models.SellerAddress{
						Street:  "Rua Jeronimo da Veiga",
						Number:  116,
						ZipCode: "04812190",
					},
				},
			},
			AcquirerID: 1,
		},
	}

	theErr := models.Error{}
	service.EXPECT().ProcessPayment(theTest.Input, &theErr).Return(theTest.ExpectedOutput)
	body, err := json.Marshal(theTest.Input)
	assert.NoError(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/process/payment", bytes.NewReader(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, theTest.StatusCode, w.Code)
	resp := w.Result()
	respBody, _ := ioutil.ReadAll(resp.Body)

	responseStruct := models.Response{}

	err = json.Unmarshal(respBody, &responseStruct)

	assert.NoError(t, err)

	if assert.Equal(t, theTest.ExpectedError, responseStruct.Errors) {
		assert.Equal(t, theTest.ExpectedOutput, responseStruct.Success)
	}
}
