package access_token

import (
	"github.com/voicurobert/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetByID(id string) (*AccessToken, *errors.RestError) {
	accessTokenId := strings.TrimSpace(id)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	return s.repository.GetByID(id)
}

func (s *service) Create(at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
