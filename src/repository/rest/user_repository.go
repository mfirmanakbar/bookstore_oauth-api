package rest

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/mfirmanakbar/bookstore_oauth-api/src/domain/users"
	"github.com/mfirmanakbar/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8182",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUserRepository {
	return &usersRepository{}
}

func (r usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
