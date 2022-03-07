package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type OrganisationRepositoryInterface interface {
	CreateOrganisation(models.Organisation, models.UserOrganisationMapping) error
}

type OrganisationRepository struct {
	db                                *gorm.DB
	UserOrganisationMappingRepository UserOrganisationMappingRepositoryInterface
}

func NewOrganisationRepository() OrganisationRepositoryInterface {
	or := new(OrganisationRepository)
	or.db = utilities.GetDatabase()
	or.UserOrganisationMappingRepository = NewUserOrganisationMappingRepository()
	return or
}

func (or *OrganisationRepository) CreateOrganisation(
	organisation models.Organisation,
	userOrganisationMapping models.UserOrganisationMapping,
) (err error) {
	db := or.db

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&organisation).Error; err != nil {
			return err
		}

		if or.UserOrganisationMappingRepository.CreateUserOrganisationMapping(userOrganisationMapping, tx); err != nil {
			return err
		}

		return nil
	})

	return err
}
