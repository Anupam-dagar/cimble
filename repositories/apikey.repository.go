package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ApiKeysRepositoryInterface interface {
	CreateApiKey(*models.ApiKey) error
	DeleteApiKey(string, string) error
	GetApiKeys(string, int64, int64) ([]models.ApiKey, int64, error)
	ValidateApiKey(string, []byte) (bool, error)
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

func (akr *ApiKeysRepository) GetApiKeys(organisationId string, offset int64, limit int64) (apiKeys []models.ApiKey, count int64, err error) {
	db := akr.Db

	db = db.Table("api_keys")
	db.Where("organisation_id = ?", organisationId)
	db.Where("revoked = false")
	db.Count(&count)
	db.Offset(int(offset))
	db.Limit(int(limit))
	err = db.Find(&apiKeys).Error

	return apiKeys, count, err
}

func (akr *ApiKeysRepository) ValidateApiKey(organisationId string, apiKey []byte) (isValid bool, err error) {
	db := akr.Db

	var count int64

	db = db.Table("api_keys")
	db.Where("organisation_id = ?", organisationId)
	db.Where("key_hash = ?", apiKey)
	db.Where("revoked = false")
	err = db.Count(&count).Error

	isValid = count > 0
	return isValid, err
}
