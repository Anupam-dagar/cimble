package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserPrivilegeRepositoryInterface interface {
	CreateUserPrivilege(*models.UserPrivilege, *gorm.DB) error
}

type UserPrivilegeRepository struct {
	db *gorm.DB
}

func NewUserPrivilegeRepository() UserPrivilegeRepositoryInterface {
	upr := new(UserPrivilegeRepository)
	upr.db = utilities.GetDatabase()
	return upr
}

func (upr *UserPrivilegeRepository) CreateUserPrivilege(userPrivilege *models.UserPrivilege, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = upr.db
	}

	err = tx.Create(userPrivilege).Error

	return err
}
