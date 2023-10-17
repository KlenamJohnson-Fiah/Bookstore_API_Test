package bin

import (
	"bookstore_oauth-api/src/domain/users"
	errors "bookstore_oauth-api/src/utils/errors_utils"
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

//
//"https://api.bookstore.com"
var (
	UsersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
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
	response := UsersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restClient response when trying to login")

	}
	if response.StatusCode > 299 {
		//fmt.Println(response.String())
		var restErr errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
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
