package services

import (
	"cimble/models"
	"cimble/repositories"
	"fmt"
)

type ConfigurationServiceInterface interface {
	CreateConfiguration(models.ConfigurationCreateRequest, string, string) (models.Configuration, error)
}

type ConfigurationService struct {
	ConfigurationRepository repositories.ConfigurationRepositoryInterface
}

func NewConfigurationService() ConfigurationServiceInterface {
	cs := new(ConfigurationService)
	cs.ConfigurationRepository = repositories.NewConfigurationRepository()
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
