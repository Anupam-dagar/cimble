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
	UpdateOrganisation(*gin.Context)
	GetOrganisations(*gin.Context)
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
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	organisation, err := oc.OrganisationService.CreateOrganisation(organisationCreatePayload, userId)

	if err != nil {
		utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, organisation)
}

func (oc *OrganisationController) UpdateOrganisation(ctx *gin.Context) {
	var organisationUpdatePayload models.OrganisationUpdateRequest
	userId := ctx.GetString("id")
	organisationId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&organisationUpdatePayload)
	if err != nil {
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	organisation, err := oc.OrganisationService.UpdateOrganisation(organisationUpdatePayload, organisationId, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusAccepted, organisation)
}

func (oc *OrganisationController) GetOrganisations(ctx *gin.Context) {
	userId := ctx.GetString("id")

	organisations, err := oc.OrganisationService.GetOrganisations(userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, organisations)
}
