package http

import (
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/bookstore_oauth-api/src/domain/access_token"
	"github.com/voicurobert/bookstore_oauth-api/src/utils/errors"
	"net/http"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (handler *accessTokenHandler) GetByID(context *gin.Context) {
	accessTokenId := context.Param("access_token_id")
	accessToken, err := handler.service.GetByID(accessTokenId)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	context.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBind(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBind(&at); err != nil {

	}
}
