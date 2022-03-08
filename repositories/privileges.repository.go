package repositories

import (
	"cimble/utilities"

	"gorm.io/gorm"
)

type PrivilegesRepositoryInterface interface {
}

type PrivilegesRepository struct {
	db *gorm.DB
}

func NewPrivilegesRepository() PrivilegesRepositoryInterface {
	pr := new(PrivilegesRepository)
	pr.db = utilities.GetDatabase()
	return pr
}
