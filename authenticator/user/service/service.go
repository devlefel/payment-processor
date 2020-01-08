package service

import (
	"authenticator/user"
	"authenticator/user/models"
)

type service struct {
	repo user.Repository
}

func NewService(r user.Repository) user.Service {
	return &service{
		repo: r,
	}
}

func (s *service) Login(login models.Login) ([]string, *models.Error) {
	userID, errs := s.repo.AuthUser(login)

	if len(errs.Internal) > 0 || len(errs.Validation) > 0 {
		return nil, errs
	}

	tokens, errs := s.repo.GetTokens(userID)

	if len(errs.Internal) > 0 || len(errs.Validation) > 0 {
		return nil, errs
	}

	return tokens, nil
}
