package rest

import (
	"bookstore_oauth-api/src/domain/users"
	errors "bookstore_oauth-api/src/utils/errors_utils"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

type usersRepository struct{}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post("/users/login")

	if err != nil {
		fmt.Println(resp.RawResponse)
		return nil, errors.NewInternalServerError("invalid restClient response when trying to login")
	}
	if resp.StatusCode() > 299 {
		//fmt.Println(response.String())
		var restErr errors.RestErr
		if err := json.Unmarshal(resp.Body(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil

}
