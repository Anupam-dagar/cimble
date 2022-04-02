package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
	"fmt"
)

type ConfigurationServiceInterface interface {
	CreateConfiguration(models.ConfigurationCreateRequest, string, string) (models.Configuration, error)
	UpdateConfiguration(models.ConfigurationUpdateRequest, string, string, string) (models.Configuration, error)
	GetConfigurations(string, string, int64, int64) (models.ConfigurationsResponse, error)
	GetFormattedConfigurations(string, string) (map[string]models.Configuration, error)
	DeleteConfiguration(string, string, string) error
}

type ConfigurationService struct {
	ConfigurationRepository repositories.ConfigurationRepositoryInterface
	UserMappingRepository   repositories.UserMappingRepositoryInterface
}

func NewConfigurationService() ConfigurationServiceInterface {
	cs := new(ConfigurationService)
	cs.ConfigurationRepository = repositories.NewConfigurationRepository()
	cs.UserMappingRepository = repositories.NewUserMappingRepository()
	return cs
}

func (cs *ConfigurationService) CreateConfiguration(
	createConfigurationPayload models.ConfigurationCreateRequest,
	createdBy string,
	projectId string,
) (configuration models.Configuration, err error) {
	configuration = createConfigurationPayload.CreateConfigurationEntity(createdBy, projectId)

	err = cs.ConfigurationRepository.CreateConfiguration(&configuration)
	if err != nil {
		fmt.Printf("error creating organisation: %v", err)
		return configuration, err
	}

	return configuration, err
}

func (cs *ConfigurationService) UpdateConfiguration(
	updateConfigurationPayload models.ConfigurationUpdateRequest,
	projectId string,
	configurationId string,
	updatedBy string,
) (configuration models.Configuration, err error) {
	configuration = updateConfigurationPayload.CreateUpdateConfigurationEntity(updatedBy)
	userProjectPrivilege, err := cs.UserMappingRepository.GetUserLevelMapping(updatedBy, projectId, constants.PROJECT)

	if err != nil {
		return configuration, err
	}

	if !userProjectPrivilege.IsUpdate {
		return configuration, fmt.Errorf(string(constants.Unauthorised))
	}

	err = cs.ConfigurationRepository.UpdateConfigurationById(&configuration, configurationId)
	if err != nil {
		fmt.Printf("error updating configuration %s by %s: %v", configuration.ID, updatedBy, err)
		return configuration, err
	}

	configuration.ID = configurationId

	return configuration, err
}

func (cs *ConfigurationService) GetConfigurations(projectId string, userId string, offset int64, limit int64) (configurations models.ConfigurationsResponse, err error) {
	userProjectPrivilege, err := cs.UserMappingRepository.GetUserLevelMapping(userId, projectId, constants.PROJECT)

	if err != nil {
		return configurations, err
	}

	if !userProjectPrivilege.IsRead {
		return configurations, fmt.Errorf(string(constants.Unauthorised))
	}

	configurationsData, count, err := cs.ConfigurationRepository.GetConfigurations(projectId, offset, limit)
	if err != nil {
		return configurations, err
	}

	page := utilities.GeneratePage(count, offset, limit)

	configurations = models.ConfigurationsResponse{
		Configurations: configurationsData,
		Page:           page,
	}

	return configurations, err
}

func (cs *ConfigurationService) GetFormattedConfigurations(
	projectId string,
	userId string,
) (configurations map[string]models.Configuration, err error) {
	configurationsData, err := cs.ConfigurationRepository.GetAllConfigurations(projectId)
	if err != nil {
		fmt.Printf("error getting formatted configurations data: %s", err.Error())
		return configurations, err
	}

	configurations = utilities.FormatConfigurations(configurationsData)
	return configurations, err
}

func (cs *ConfigurationService) DeleteConfiguration(
	projectId string,
	configurationId string,
	deletedBy string,
) (err error) {
	userProjectPrivilege, err := cs.UserMappingRepository.GetUserLevelMapping(deletedBy, projectId, constants.PROJECT)

	if err != nil {
		return err
	}

	if !userProjectPrivilege.IsDelete {
		return fmt.Errorf(string(constants.Unauthorised))
	}

	err = cs.ConfigurationRepository.DeleteConfigurationById(configurationId, deletedBy)
	if err != nil {
		fmt.Printf("error deleting configuration %s by %s: %v", configurationId, deletedBy, err)
		return err
	}

	return err
}
