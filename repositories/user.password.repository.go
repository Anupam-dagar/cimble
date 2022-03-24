package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserPasswordRepositoryInterface interface {
	CreateUserPassword(models.UserPassword, *gorm.DB) error
}

type UserPasswordRepository struct {
	db *gorm.DB
}

func NewUserPasswordRepository() UserPasswordRepositoryInterface {
	upr := new(UserPasswordRepository)
	upr.db = utilities.GetDatabase()
	return upr
}

func (upr *UserPasswordRepository) CreateUserPassword(userPassword models.UserPassword, tx *gorm.DB) error {
	if tx == nil {
		tx = upr.db
	}

	err := tx.Create(&userPassword).Error

	return err
}
