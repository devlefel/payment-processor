package user

import "authenticator/user/models"

type Service interface {
	Login(login models.Login) ([]string, *models.Error)
}
