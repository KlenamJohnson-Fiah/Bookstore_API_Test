package access_token

import (
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/domain/users"
	"bookstore_oauth-api/src/repository/db"
	"bookstore_oauth-api/src/repository/rest"
	errors "bookstore_oauth-api/src/utils/errors_utils"
)

// type Repository interface {
// 	GetById(string) (*access_token.AccessToken, *errors.RestErr)
// 	Create(access_token.AccessToken) (*access_token.AccessToken, *errors.RestErr)
// 	UpdateExpirationTime(*access_token.AccessToken) *errors.RestErr
// }

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(at access_token.AccessToken) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(at *access_token.AccessToken) *errors.RestErr
}

type service struct {
	userRepository   rest.RestUsersRepository
	accessRepository db.DbRepository
}

func NewService(repo rest.RestUsersRepository, accessTokenRepo db.DbRepository) Service {
	return &service{
		userRepository:   repo,
		accessRepository: accessTokenRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.accessRepository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request users.UserLoginRequest) (*access_token.AccessToken, *errors.RestErr) {
	// Authenticate the user against the Users API:
	user, err := s.userRepository.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.accessRepository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at *access_token.AccessToken) *errors.RestErr {
	if err := at.ValidateAccessToken(); err != nil {
		return err
	}

	return s.UpdateExpirationTime(at)
}
