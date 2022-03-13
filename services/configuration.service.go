package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"fmt"
)

type ConfigurationServiceInterface interface {
	CreateConfiguration(models.ConfigurationCreateRequest, string, string) (models.Configuration, error)
	UpdateConfiguration(models.ConfigurationUpdateRequest, string, string, string) (models.Configuration, error)
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

	return configuration, err
}
