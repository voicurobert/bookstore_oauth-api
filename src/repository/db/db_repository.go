package db

import (
	"github.com/gocql/gocql"
	"github.com/voicurobert/bookstore_oauth-api/src/clients/cassandra"
	"github.com/voicurobert/bookstore_oauth-api/src/domain/access_token"
	"github.com/voicurobert/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdate            = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {
	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (db *dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().
		Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).
		Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().
		Query(queryUpdate, at.Expires, at.AccessToken).
		Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
