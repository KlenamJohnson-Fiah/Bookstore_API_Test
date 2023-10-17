package access_token

import (
	errors "bookstore_oauth-api/src/utils/errors_utils"
	"fmt"

	"strings"
	"time"

	"bookstore_oauth-api/src/utils/crypto"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//used for grant_type
	Username string `json:"Username"`
	Password string `json:"password"`

	//used for client credentials
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) ValidateAccessToken() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) ValidateAccessToken() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client ID")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) *AccessToken {
	return &AccessToken{
		UserID:  userId,
		Expires: time.Now().UTC().Add(time.Hour * expirationTime).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
