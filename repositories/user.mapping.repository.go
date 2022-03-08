package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserMappingRepositoryInterface interface {
	CreateUserMapping(*models.UserMapping, *gorm.DB) error
}

type UserMappingRepository struct {
	db *gorm.DB
}

func NewUserMappingRepository() UserMappingRepositoryInterface {
	upr := new(UserMappingRepository)
	upr.db = utilities.GetDatabase()
	return upr
}

func (upr *UserMappingRepository) CreateUserMapping(userMapping *models.UserMapping, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = upr.db
	}

	err = tx.Create(userMapping).Error

	return err
}
