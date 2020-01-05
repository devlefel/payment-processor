package payment

import "processor/payment/models"

type Repository interface {
	GetCardSensitiveData(token string, errors *models.Error) *models.CardSensitiveData
	GetAcquirerURL(id int64, errors *models.Error) string
}
