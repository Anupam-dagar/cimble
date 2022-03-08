package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserOrganisationMappingRepositoryInterface interface {
	CreateUserOrganisationMapping(*models.UserOrganisationMapping, *gorm.DB) error
}

type UserOrganisationMappingRepository struct {
	db *gorm.DB
}

func NewUserOrganisationMappingRepository() UserOrganisationMappingRepositoryInterface {
	uomr := new(UserOrganisationMappingRepository)
	uomr.db = utilities.GetDatabase()
	return uomr
}

func (uomr *UserOrganisationMappingRepository) CreateUserOrganisationMapping(
	userOrganisationMapping *models.UserOrganisationMapping,
	tx *gorm.DB,
) (err error) {
	if tx == nil {
		tx = uomr.db
	}

	err = tx.Create(userOrganisationMapping).Error

	return err
}
