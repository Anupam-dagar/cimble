package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	CreateProject(*models.Project, *models.UserMapping) error
	UpdateProjectById(*models.Project, string) error
	GetProjects(string, string, int64, int64) ([]models.ProjectModel, int64, error)
	DeleteProjectById(string, string) error
	DeleteProjectByOrganisationId(*gorm.DB, string, string) error
}

type ProjectRepository struct {
	db                      *gorm.DB
	UserMappingRepository   UserMappingRepositoryInterface
	ConfigurationRepository ConfigurationRepositoryInterface
}

func NewProjectRepository() ProjectRepositoryInterface {
	pr := new(ProjectRepository)
	pr.db = utilities.GetDatabase()
	pr.UserMappingRepository = NewUserMappingRepository()
	pr.ConfigurationRepository = NewConfigurationRepository()
	return pr
}

func (pr *ProjectRepository) CreateProject(
	project *models.Project,
	userMapping *models.UserMapping,
) (err error) {
	db := pr.db
	db = db.Table("projects")

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
	db = db.Table("projects")

	err = db.Where("id = ?", projectId).Updates(project).Error

	return err
}

func (pr *ProjectRepository) GetProjects(userId string, organisationId string, offset int64, limit int64) (projects []models.ProjectModel, count int64, err error) {
	db := pr.db

	db = db.Table("projects")
	db.Select("projects.*, count(configurations.id) as configurations_count")
	db.Joins("inner join user_mappings on user_mappings.level_id = projects.id and user_mappings.user_id = ?", userId)
	db.Joins("left join configurations on projects.id = configurations.project_id")
	db.Where("projects.organisation_id = ?", organisationId)
	db.Count(&count)
	db.Offset(int(offset))
	db.Limit(int(limit))
	db.Group("projects.id")
	err = db.Find(&projects).Error

	return projects, count, err
}

func (pr *ProjectRepository) DeleteProjectById(projectId string, deletedBy string) (err error) {
	db := pr.db
	db = db.Table("projects")

	var project models.Project
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", projectId).Find(&project).Error; err != nil {
			return err
		}

		if err := tx.Delete(&project).Error; err != nil {
			return err
		}

		projectArchive := project.CreateProjectArchiveEntity(deletedBy)
		if err := tx.Table("project_archives").Create(&projectArchive).Error; err != nil {
			return err
		}

		if err := pr.ConfigurationRepository.DeleteConfigurationByProjectId(tx, projectId, deletedBy); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (pr *ProjectRepository) DeleteProjectByOrganisationId(tx *gorm.DB, organisationId string, deletedBy string) (err error) {
	if tx == nil {
		tx = pr.db
	}

	tx = tx.Table("projects")

	var projects []models.Project
	err = tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("organisation_id = ?", organisationId).Find(&projects).Error; err != nil {
			return err
		}

		if len(projects) == 0 {
			return nil
		}

		if err := tx.Delete(&projects).Error; err != nil {
			return err
		}

		var projectsArchive []models.ProjectArchive
		var projectIds []string
		for _, deletedProject := range projects {
			projectsArchive = append(projectsArchive, deletedProject.CreateProjectArchiveEntity(deletedBy))
			projectIds = append(projectIds, deletedProject.ID)
		}
		if err := tx.Table("project_archives").Create(&projectsArchive).Error; err != nil {
			return err
		}

		if err := pr.ConfigurationRepository.DeleteConfigurationByProjectIds(tx, projectIds, deletedBy); err != nil {
			return err
		}

		return nil
	})

	return err
}
