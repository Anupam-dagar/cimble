package controllers

import (
	"cimble/models"
	"cimble/services"
	"cimble/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectControllerInterface interface {
	CreateProject(ctx *gin.Context)
}

type ProjectController struct {
	ProjectService services.ProjectServiceInterface
}

func NewProjectController() ProjectControllerInterface {
	pc := new(ProjectController)
	pc.ProjectService = services.NewProjectService()
	return pc
}

func (pc *ProjectController) CreateProject(ctx *gin.Context) {
	var createProjectPayload models.ProjectCreateRequest
	userId := ctx.GetString("id")

	err := ctx.ShouldBindJSON(&createProjectPayload)
	if err != nil {
		utilities.ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	project, err := pc.ProjectService.CreateProject(createProjectPayload, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, project)
}
