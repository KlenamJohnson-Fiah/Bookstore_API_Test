package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/go-resty/resty/v2"
// 	"github.com/jarcoal/httpmock"
// 	"github.com/stretchr/testify/assert"
// 	"bookstore_oauth-api/src/utils/errors_utils"
// )

// const (
// 	BaseURL = "localhost:8080//users/login"
// )

// func TestLoginUserTimeoutFromApi(t *testing.T) {
// 	client := resty.New()
// 	httpmock.ActivateNonDefault(client.GetClient())

// 	// userStr := users.UserLoginRequest{
// 	// 	Email: "email@test.com",
// 	// 	Password: "password",
// 	// }

// 	httpmock.RegisterResponder("POST",BaseURL,func(req *http.Request) (*http.Response, error) {
// 		return httpmock.NewJsonResponse(400, `{
// 				"message": "invalid login credentials",
// 			 	"status": "400",
// 			 	"error": "not_found",
// 			 }`)
// 	}
// 	testUser := NewRepository()
// 	client.R().Post(BaseURL)
// 	resp, err  := testUser.LoginUser("email@test.com", "password")
// 	//assert.Equal(t, `"email": "email@test.com", "password": "password"`, resp)
// 	//assert.Equal(t, "invalid restClient response when trying to login", err.Message)
// 	assert.EqualValues(t, http.StatusInternalServerError,err.Status)
// 	assert.NotNil(t, "email@test.com", resp)
// 	fmt.Println(resp)
// 	fmt.Println(err)
// 	httpmock.DeactivateAndReset()

// }
