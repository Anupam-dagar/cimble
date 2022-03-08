package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	CreateProject(*models.Project, *models.UserMapping) error
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
