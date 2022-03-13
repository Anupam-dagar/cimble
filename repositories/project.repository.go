package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	CreateProject(*models.Project, *models.UserMapping) error
	UpdateProjectById(*models.Project, string) error
	GetProjects(string) ([]models.Project, error)
}

type ProjectRepository struct {
	db                    *gorm.DB
	UserMappingRepository UserMappingRepositoryInterface
}

func NewProjectRepository() ProjectRepositoryInterface {
	pr := new(ProjectRepository)
	pr.db = utilities.GetDatabase()
	pr.UserMappingRepository = NewUserMappingRepository()
	return pr
}

func (pr *ProjectRepository) CreateProject(
	project *models.Project,
	userMapping *models.UserMapping,
) (err error) {
	db := pr.db

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(project).Error; err != nil {
			return err
		}

		if pr.UserMappingRepository.CreateUserMapping(userMapping, tx); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (pr *ProjectRepository) UpdateProjectById(project *models.Project, projectId string) (err error) {
	db := pr.db

	err = db.Where("id = ?", projectId).Updates(project).Error

	return err
}

func (pr *ProjectRepository) GetProjects(userId string) (projects []models.Project, err error) {
	db := pr.db

	db.Select("projects.*")
	db.Joins("inner join user_mappings on user_mappings.level_id = projects.id and user_mappings.user_id = ?", userId)
	err = db.Find(&projects).Error

	return projects, err
}
