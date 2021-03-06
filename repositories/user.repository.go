package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(models.User, models.UserPassword) error
	GetUserById(userId string) (models.User, error)
	GetUserWithPassword(string) (models.User, error)
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

func (ur *UserRepository) GetUserById(userId string) (user models.User, err error) {
	db := ur.db

	err = db.Find(&user, "id = ?", userId).Error

	return user, err
}

func (ur *UserRepository) GetUserWithPassword(userEmail string) (models.User, error) {
	db := ur.db

	var user models.User

	db = db.Table("users")
	db.Preload("UserPassword")
	db.Where("users.email = ?", userEmail)
	err := db.Find(&user).Error
	return user, err
}
