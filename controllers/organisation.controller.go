package controllers

import (
	"cimble/models"
	"cimble/services"
	"cimble/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganisationControllerInterface interface {
	CreateOrganisation(*gin.Context)
}

type OrganisationController struct {
	OrganisationService services.OrganisationServiceInterface
}

func NewOrganisationController() OrganisationControllerInterface {
	oc := new(OrganisationController)
	oc.OrganisationService = services.NewOrganisationService()
	return oc
}

func (oc *OrganisationController) CreateOrganisation(ctx *gin.Context) {
	var organisationCreatePayload models.OrganisationCreateRequest
	userId := ctx.GetString("id")

	err := ctx.ShouldBindJSON(&organisationCreatePayload)
	if err != nil {
		utilities.ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	organisation, err := oc.OrganisationService.CreateOrganisation(organisationCreatePayload, userId)

	if err != nil {
		utilities.ResponseWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, organisation)
}
