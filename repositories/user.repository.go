package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepositoryInterface {
	ur := new(UserRepository)
	ur.db = utilities.GetDatabase()
	return ur
}

func (ur *UserRepository) AddUser(user models.User) error {
	db := ur.db

	user.ID = ksuid.New().String()
	user.UpdatedBy = user.CreatedBy
	result := db.Create(&user)

	return result.Error
}
