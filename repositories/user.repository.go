package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(models.User, models.UserPassword) error
}

type UserRepository struct {
	db                     *gorm.DB
	UserPasswordRepository UserPasswordRepositoryInterface
}

func NewUserRepository() UserRepositoryInterface {
	ur := new(UserRepository)
	ur.db = utilities.GetDatabase()
	ur.UserPasswordRepository = NewUserPasswordRepository()
	return ur
}

func (ur *UserRepository) AddUser(user models.User, userPassword models.UserPassword) error {
	db := ur.db

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := ur.UserPasswordRepository.CreateUserPassword(userPassword, tx); err != nil {
			return err
		}

		return nil
	})

	return err
}
