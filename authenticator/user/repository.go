package user

import "authenticator/user/models"

type Repository interface {
	AuthUser(login models.Login) (int64, *models.Error)
	GetTokens(userID int64) ([]string, *models.Error)
}
