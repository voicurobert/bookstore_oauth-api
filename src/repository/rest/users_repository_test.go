package rest

import "testing"

func TestLoginUserTimeoutFromApi(t *testing.T) {
	repository := usersRepository{}
	_, _ = repository.LoginUser("email@gmail.com", "password")
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
