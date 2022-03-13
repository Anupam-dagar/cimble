package controllers

import (
	"cimble/models"
	"cimble/services"
	"cimble/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigurationControllerInterface interface {
	CreateConfiguration(*gin.Context)
}

type ConfigurationController struct {
	ConfigurationService services.ConfigurationServiceInterface
}

func NewConfigurationController() ConfigurationControllerInterface {
	cc := new(ConfigurationController)
	cc.ConfigurationService = services.NewConfigurationService()
	return cc
}

func (cc *ConfigurationController) CreateConfiguration(ctx *gin.Context) {
	var createConfigurationPayload models.ConfigurationCreateRequest
	userId := ctx.GetString("id")
	projectId := ctx.Param("projectId")

	err := utilities.GetRequestBody(ctx, &createConfigurationPayload)
	if err != nil {
		return
	}

	configuration, err := cc.ConfigurationService.CreateConfiguration(createConfigurationPayload, userId, projectId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, configuration)
}
