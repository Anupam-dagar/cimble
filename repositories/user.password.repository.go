package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserPasswordRepositoryInterface interface {
	CreateUserPassword(models.UserPassword, *gorm.DB) error
	GetUserPassword(string) (models.UserJoinUserPassword, error)
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

func (upr *UserPasswordRepository) GetUserPassword(userEmail string) (models.UserJoinUserPassword, error) {
	db := upr.db

	var userJoinUserPassword models.UserJoinUserPassword

	db = db.Table("user_passwords")
	db.Select("users.email, users.id, user_passwords.*")
	db.Joins("inner join users on users.id = user_passwords.user_id")
	db.Where("users.email = ?", userEmail)
	err := db.Find(&userJoinUserPassword).Error

	return userJoinUserPassword, err
}
