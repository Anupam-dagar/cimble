package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
	"fmt"
)

type ProjectServiceInterface interface {
	CreateProject(createProjectPayload models.ProjectCreateRequest, createdBy string) (project models.Project, err error)
	UpdateProject(models.ProjectUpdateRequest, string, string) (models.Project, error)
	GetProjects(string, string, int64, int64) (models.ProjectsResponse, error)
	DeleteProject(string, string) error
}

type ProjectService struct {
	ProjectRepository     repositories.ProjectRepositoryInterface
	UserMappingRepository repositories.UserMappingRepositoryInterface
}

func NewProjectService() ProjectServiceInterface {
	ps := new(ProjectService)
	ps.ProjectRepository = repositories.NewProjectRepository()
	ps.UserMappingRepository = repositories.NewUserMappingRepository()
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

func (ps *ProjectService) UpdateProject(
	projectPayload models.ProjectUpdateRequest,
	projectId,
	updatedBy string,
) (project models.Project, err error) {
	project = projectPayload.CreateUpdateProjectEntity(updatedBy)
	userProjectPrivilege, err := ps.UserMappingRepository.GetUserLevelMapping(updatedBy, projectId, constants.PROJECT)

	if err != nil {
		return project, err
	}

	if !userProjectPrivilege.IsUpdate {
		return project, fmt.Errorf(string(constants.Unauthorised))
	}

	err = ps.ProjectRepository.UpdateProjectById(&project, projectId)
	if err != nil {
		fmt.Printf("error updating project %s by %s: %v", projectId, updatedBy, err)
		return project, err
	}

	project.ID = projectId

	return project, err
}

func (ps *ProjectService) GetProjects(userId string, organisationId string, offset int64, limit int64) (projects models.ProjectsResponse, err error) {
	projectsData, count, err := ps.ProjectRepository.GetProjects(userId, organisationId, offset, limit)
	if err != nil {
		return projects, err
	}

	page := utilities.GeneratePage(count, offset, limit)

	projects = models.ProjectsResponse{
		Projects: projectsData,
		Page:     page,
	}
	return projects, err
}

func (ps *ProjectService) DeleteProject(
	projectId,
	deletedBy string,
) (err error) {
	userProjectPrivilege, err := ps.UserMappingRepository.GetUserLevelMapping(deletedBy, projectId, constants.PROJECT)

	if err != nil {
		return err
	}

	if !userProjectPrivilege.IsDelete {
		return fmt.Errorf(string(constants.Unauthorised))
	}

	err = ps.ProjectRepository.DeleteProjectById(projectId, deletedBy)
	if err != nil {
		fmt.Printf("error deleting project %s by %s: %v", projectId, deletedBy, err)
		return err
	}

	return err
}
