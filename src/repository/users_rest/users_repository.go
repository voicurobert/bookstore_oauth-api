package users_rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/voicurobert/bookstore_oauth-api/src/domain/users"
	"github.com/voicurobert/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	restClient rest.RequestBuilder
)

func init() {
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
}

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestError) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := restClient.Post("/users/login", req)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid users_rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
