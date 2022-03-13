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
	UpdateConfiguration(ctx *gin.Context)
	GetConfigurations(ctx *gin.Context)
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

func (cc *ConfigurationController) UpdateConfiguration(ctx *gin.Context) {
	var updateConfigurationPayload models.ConfigurationUpdateRequest
	userId := ctx.GetString("id")
	projectId := ctx.Param("projectId")
	configurationId := ctx.Param("id")

	err := utilities.GetRequestBody(ctx, &updateConfigurationPayload)
	if err != nil {
		return
	}

	project, err := cc.ConfigurationService.UpdateConfiguration(updateConfigurationPayload, projectId, configurationId, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusAccepted, project)
}

func (cc *ConfigurationController) GetConfigurations(ctx *gin.Context) {
	userId := ctx.GetString("id")
	projectId := ctx.Param("projectId")

	configurations, err := cc.ConfigurationService.GetConfigurations(projectId, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, configurations)
}
