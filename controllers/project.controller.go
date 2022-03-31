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
	UpdateProject(ctx *gin.Context)
	GetProjects(*gin.Context)
	DeleteProject(ctx *gin.Context)
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
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	project, err := pc.ProjectService.CreateProject(createProjectPayload, userId)
	if err != nil {
		utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusCreated, project)
}

func (pc *ProjectController) UpdateProject(ctx *gin.Context) {
	var updateProjectPayload models.ProjectUpdateRequest
	userId := ctx.GetString("id")
	projectId := ctx.Param("id")

	err := utilities.GetRequestBody(ctx, &updateProjectPayload)
	if err != nil {
		return
	}

	project, err := pc.ProjectService.UpdateProject(updateProjectPayload, projectId, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusAccepted, project)
}

func (pc *ProjectController) GetProjects(ctx *gin.Context) {
	userId := ctx.GetString("id")
	organisationId := ctx.Param("organisationId")

	projects, err := pc.ProjectService.GetProjects(userId, organisationId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, projects)
}

func (pc *ProjectController) DeleteProject(ctx *gin.Context) {
	userId := ctx.GetString("id")
	projectId := ctx.Param("id")

	err := pc.ProjectService.DeleteProject(projectId, userId)
	if err != nil {
		utilities.ResponseWithError(ctx, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusAccepted, "")
}
