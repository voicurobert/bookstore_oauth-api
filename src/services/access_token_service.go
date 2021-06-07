package services

import (
	"github.com/voicurobert/bookstore_oauth-api/src/domain/access_token"
	"github.com/voicurobert/bookstore_oauth-api/src/repository/db"
	"github.com/voicurobert/bookstore_oauth-api/src/repository/users_rest"
	"github.com/voicurobert/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type service struct {
	repository users_rest.RestUsersRepository
	dbRepo     db.DBRepository
}

func NewService(repo users_rest.RestUsersRepository, dbRepo db.DBRepository) Service {
	return &service{repository: repo, dbRepo: dbRepo}
}

func (s *service) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {
	accessTokenId := strings.TrimSpace(id)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	return s.dbRepo.GetByID(accessTokenId)
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	// TODO support both client_credentials and password grant_types
	user, err := s.repository.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
