package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ApiKeysRepositoryInterface interface {
	CreateApiKey(*models.ApiKey) error
}

type ApiKeysRepository struct {
	Db *gorm.DB
}

func NewApiKeysRepository() ApiKeysRepositoryInterface {
	akr := new(ApiKeysRepository)
	akr.Db = utilities.GetDatabase()
	return akr
}

func (akr *ApiKeysRepository) CreateApiKey(apiKey *models.ApiKey) error {
	db := akr.Db

	db = db.Table("api_keys")
	err := db.Create(apiKey).Error

	return err
}
