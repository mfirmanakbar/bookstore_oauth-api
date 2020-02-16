package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/mfirmanakbar/bookstore_oauth-api/src/domain/users"
	"github.com/mfirmanakbar/bookstore_utils-go/rest_errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8182",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUserRepository {
	return &usersRepository{}
}

func (r usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError(
			"invalid restclient response when trying to login user",
			errors.New("database error"),
		)
	}

	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError(
				"invalid error interface when trying to login user",
				errors.New("database error"),
			)
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError(
			"error when trying to unmarshal users login response",
			errors.New("database error"),
		)
	}
	return &user, nil
}
