package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type OrganisationRepositoryInterface interface {
	CreateOrganisation(*models.Organisation, *models.UserMapping) error
	UpdateOrganisationById(*models.Organisation, string) error
	GetOrganisations(string) ([]models.Organisation, error)
}

type OrganisationRepository struct {
	db                                *gorm.DB
	UserOrganisationMappingRepository UserOrganisationMappingRepositoryInterface
	UserMappingRepository             UserMappingRepositoryInterface
}

func NewOrganisationRepository() OrganisationRepositoryInterface {
	or := new(OrganisationRepository)
	or.db = utilities.GetDatabase()
	or.UserOrganisationMappingRepository = NewUserOrganisationMappingRepository()
	or.UserMappingRepository = NewUserMappingRepository()
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

func (or *OrganisationRepository) GetOrganisations(userId string) (organisations []models.Organisation, err error) {
	db := or.db

	db.Select("organisations.*")
	db.Joins("inner join user_mappings on user_mappings.level_id = organisations.id and user_mappings.user_id = ?", userId)
	err = db.Find(&organisations).Error

	return organisations, err
}
