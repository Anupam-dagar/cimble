package repositories

import (
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type ConfigurationRepositoryInterface interface {
	CreateConfiguration(*models.Configuration) error
	UpdateConfigurationById(*models.Configuration, string) error
	GetConfigurations(string) ([]models.Configuration, error)
}

type ConfigurationRepository struct {
	Db *gorm.DB
}

func NewConfigurationRepository() ConfigurationRepositoryInterface {
	cr := new(ConfigurationRepository)
	cr.Db = utilities.GetDatabase()
	return cr
}

func (cr *ConfigurationRepository) CreateConfiguration(
	configuration *models.Configuration,
) (err error) {
	db := cr.Db

	db = db.Table("configurations")
	err = db.Create(configuration).Error

	return err
}

func (cr *ConfigurationRepository) UpdateConfigurationById(configuration *models.Configuration, configurationId string) (err error) {
	db := cr.Db

	db = db.Table("configurations")
	err = db.Where("id = ?", configurationId).Updates(configuration).Error

	return err
}

func (cr *ConfigurationRepository) GetConfigurations(projectId string) (configurations []models.Configuration, err error) {
	db := cr.Db

	db = db.Table("configurations")
	db.Where("project_id = ?", projectId)
	err = db.Find(&configurations).Error

	return configurations, err
}
