package rest

// next time try using https://github.com/vektra/mockery

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer() // start to mocking
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups() //remove all prev mocks
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8282/users/login",
		ReqBody:      `{"email":"firman@gmail.com","password":"firman123"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8282/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404_string", "error": "not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8282/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8282/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"1_string","first_name":"Firman","last_name":"Zaidin","email":"firman@gmail.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8282/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":1,"first_name":"Firman","last_name":"Zaidin","email":"firman@gmail.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Firman", user.FirstName)
	assert.EqualValues(t, "Zaidin", user.LastName)
	assert.EqualValues(t, "firman@gmail.com", user.Email)
}
