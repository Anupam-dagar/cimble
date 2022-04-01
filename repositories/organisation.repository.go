package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type OrganisationRepositoryInterface interface {
	CreateOrganisation(*models.Organisation, *models.UserMapping) error
	UpdateOrganisationById(*models.Organisation, string) error
	GetOrganisations(string, int64, int64) ([]models.OrganisationModel, int64, error)
	DeleteOrganisationById(string, string) error
}

type OrganisationRepository struct {
	db                                *gorm.DB
	UserOrganisationMappingRepository UserOrganisationMappingRepositoryInterface
	UserMappingRepository             UserMappingRepositoryInterface
	ProjectRepository                 ProjectRepositoryInterface
}

func NewOrganisationRepository() OrganisationRepositoryInterface {
	or := new(OrganisationRepository)
	or.db = utilities.GetDatabase()
	or.UserOrganisationMappingRepository = NewUserOrganisationMappingRepository()
	or.UserMappingRepository = NewUserMappingRepository()
	or.ProjectRepository = NewProjectRepository()
	return or
}

func (or *OrganisationRepository) CreateOrganisation(
	organisation *models.Organisation,
	userMapping *models.UserMapping,
) (err error) {
	db := or.db

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(organisation).Error; err != nil {
			return err
		}

		if or.UserMappingRepository.CreateUserMapping(userMapping, tx); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (or *OrganisationRepository) UpdateOrganisationById(organisation *models.Organisation, organisationId string) (err error) {
	db := or.db

	err = db.Where("id = ?", organisationId).Updates(organisation).Error

	return err
}

func (or *OrganisationRepository) GetOrganisations(userId string, offset int64, limit int64) (organisations []models.OrganisationModel, count int64, err error) {
	db := or.db

	db = db.Table("organisations")
	db.Select("organisations.*, count(projects.id) as projects_count")
	db.Joins("inner join user_mappings on user_mappings.level_id = organisations.id and user_mappings.user_id = ?", userId)
	db.Joins("left join projects on projects.organisation_id = organisations.id")
	db.Group("organisations.id")
	db.Offset(int(offset))
	db.Limit(int(limit))
	err = db.Find(&organisations).Error

	if err == nil {
		db.Count(&count)
	}

	return organisations, count, err
}

func (or *OrganisationRepository) DeleteOrganisationById(organisationId string, deletedBy string) (err error) {
	db := or.db
	db = db.Table("organisations")

	var organisation models.Organisation

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", organisationId).Find(&organisation).Error; err != nil {
			return err
		}

		if err := tx.Delete(&organisation).Error; err != nil {
			return err
		}

		if err := or.ProjectRepository.DeleteProjectByOrganisationId(tx, organisationId, deletedBy); err != nil {
			return err
		}

		organisationArchive := organisation.CreateOrganisationArchiveEntity(deletedBy)
		if err := tx.Table("organisation_archives").Create(&organisationArchive).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
