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
	DeleteConfigurationById(string, string) error
	DeleteConfigurationByProjectId(*gorm.DB, string, string) error
	DeleteConfigurationByProjectIds(*gorm.DB, []string, string) error
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

func (cr *ConfigurationRepository) DeleteConfigurationById(configurationId string, deletedBy string) (err error) {
	db := cr.Db

	var configuration models.Configuration
	db = db.Table("configurations")
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", configurationId).Find(&configuration).Error; err != nil {
			return err
		}

		if err := tx.Delete(&configuration).Error; err != nil {
			return err
		}

		configurationArchive := configuration.CreateConfigurationArchiveEntity(deletedBy)
		if err := tx.Table("configuration_archives").Create(&configurationArchive).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (cr *ConfigurationRepository) DeleteConfigurationByProjectId(tx *gorm.DB, projectId string, deletedBy string) (err error) {
	if tx == nil {
		tx = cr.Db
	}

	var configuration []models.Configuration
	tx = tx.Table("configurations")
	err = tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("project_id = ?", projectId).Find(&configuration).Error; err != nil {
			return err
		}

		if err := tx.Delete(&configuration).Error; err != nil {
			return err
		}

		var configurationArchive []models.ConfigurationArchive
		for _, deletedConfiguration := range configuration {
			configurationArchive = append(configurationArchive, deletedConfiguration.CreateConfigurationArchiveEntity(deletedBy))
		}
		if err := tx.Table("configuration_archives").Create(&configurationArchive).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (cr *ConfigurationRepository) DeleteConfigurationByProjectIds(tx *gorm.DB, projectIds []string, deletedBy string) (err error) {
	if tx == nil {
		tx = cr.Db
	}

	var configuration []models.Configuration
	tx = tx.Table("configurations")
	err = tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("project_id in (?)", projectIds).Find(&configuration).Error; err != nil {
			return err
		}

		if err := tx.Delete(&configuration).Error; err != nil {
			return err
		}

		var configurationArchive []models.ConfigurationArchive
		for _, deletedConfiguration := range configuration {
			configurationArchive = append(configurationArchive, deletedConfiguration.CreateConfigurationArchiveEntity(deletedBy))
		}
		if err := tx.Table("configuration_archives").Create(&configurationArchive).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
