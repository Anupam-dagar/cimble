package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"fmt"
)

type ProjectServiceInterface interface {
	CreateProject(createProjectPayload models.ProjectCreateRequest, createdBy string) (project models.Project, err error)
}

type ProjectService struct {
	ProjectRepository repositories.ProjectRepositoryInterface
}

func NewProjectService() ProjectServiceInterface {
	ps := new(ProjectService)
	ps.ProjectRepository = repositories.NewProjectRepository()
	return ps
}

func (ps *ProjectService) CreateProject(
	createProjectPayload models.ProjectCreateRequest,
	createdBy string,
) (project models.Project, err error) {
	project = createProjectPayload.CreateProjectEntity(createdBy)
	userMapping := models.UserMapping{
		UserId:    createdBy,
		LevelFor:  string(constants.PROJECT),
		LevelId:   project.ID,
		Privelege: string(constants.OWNER),
		BaseEntity: models.BaseEntity{
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	}

	err = ps.ProjectRepository.CreateProject(&project, &userMapping)
	if err != nil {
		fmt.Printf("error creating organisation: %v", err)
		return project, err
	}

	return project, err
}
