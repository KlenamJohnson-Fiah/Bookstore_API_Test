package http

import (
	"bookstore_oauth-api/src/domain/access_token"
	errors "bookstore_oauth-api/src/utils/errors_utils"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json Request")
		c.JSON(restErr.Status, restErr)

	}
	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)

	}
	c.JSON(http.StatusCreated, at)

}
