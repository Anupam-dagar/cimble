package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ApiKeysRepositoryInterface interface {
	CreateApiKey(*models.ApiKey) error
	DeleteApiKey(string, string) error
	GetApiKeys(string) ([]models.ApiKey, error)
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

func (akr *ApiKeysRepository) DeleteApiKey(apiKeyId string, deletedBy string) (err error) {
	db := akr.Db

	updateApiKeyEntity := models.ApiKey{
		Revoked: 1,
		BaseEntity: models.BaseEntity{
			UpdatedBy: deletedBy,
		},
	}

	db = db.Table("api_keys")
	db.Where("id = ?", apiKeyId)
	err = db.Updates(&updateApiKeyEntity).Error

	return err
}

func (akr *ApiKeysRepository) GetApiKeys(organisationId string) (apiKeys []models.ApiKey, err error) {
	db := akr.Db

	db = db.Table("api_keys")
	db.Where("organisation_id = ?", organisationId)
	db.Where("revoked = false")
	err = db.Find(&apiKeys).Error

	return apiKeys, err
}
