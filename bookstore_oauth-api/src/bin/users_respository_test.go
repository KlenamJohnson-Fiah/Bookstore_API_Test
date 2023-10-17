package bin

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start Test Cases")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8080/users/login",
		ReqBody:      `{"email":"test1@test.com", "password":"64a30b9ad662bc1b53bcccba090b44da"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test1@test.com", "64a30b9ad662bc1b53bcccba090b44da")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restClient response when trying to login", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8080/users/login",
		ReqBody:      `{"email":"test1@test.com", "password":"64a30b9ad662bc1b53bcccba090b44da"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid error interface when trying to login user", "status": "404", "error":"not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test1@test.com", "64a30b9ad662bc1b53bcccba090b44da")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	fmt.Println(err.Message)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)

}

func TestLoginUserNotFoundCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test@email.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)

}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Dinamo", "last_name": "Timpo7", "email": "test2@test.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test@email.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users response", err.Message)

}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Dinamo", "last_name": "Timpo7", "email": "test2@test.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test@email.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Dinamo", user.FirstName)
	assert.EqualValues(t, "Timpo7", user.LastName)
	assert.EqualValues(t, "test2@test.com", user.Email)

}
