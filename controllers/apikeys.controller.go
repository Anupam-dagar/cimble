package controllers

import (
	"cimble/models"
	"cimble/services"
	"cimble/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiKeysControllerInterface interface {
	CreateApiKey(*gin.Context)
}

type ApiKeysController struct {
	ApiKeyService services.ApiKeyServiceInterface
}

func NewApiKeysController() ApiKeysControllerInterface {
	ac := new(ApiKeysController)
	ac.ApiKeyService = services.NewApiKeyService()
	return ac
}

func (akc *ApiKeysController) CreateApiKey(ctx *gin.Context) {
	userId := ctx.GetString("id")

	var createApiKeyRequest models.ApiKeyCreateRequest
	err := utilities.GetRequestBody(ctx, &createApiKeyRequest)
	if err != nil {
		return
	}

	apiKey, err := akc.ApiKeyService.CreateApiKey(createApiKeyRequest, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, apiKey)
}
