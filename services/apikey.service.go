package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
	"fmt"
)

type ApiKeyServiceInterface interface {
	CreateApiKey(models.ApiKeyCreateRequest, string) (models.ApiKey, error)
	DeleteApiKey(string, string, string) error
}

type ApiKeyService struct {
	UserMappingRepository repositories.UserMappingRepositoryInterface
	ApiKeyRepository      repositories.ApiKeysRepositoryInterface
}

func NewApiKeyService() ApiKeyServiceInterface {
	aks := new(ApiKeyService)
	aks.UserMappingRepository = repositories.NewUserMappingRepository()
	aks.ApiKeyRepository = repositories.NewApiKeysRepository()
	return aks
}

func (aks *ApiKeyService) CreateApiKey(createApiKeyRequest models.ApiKeyCreateRequest, createdBy string) (apiKey models.ApiKey, err error) {
	userProjectPrivilege, err := aks.UserMappingRepository.GetUserLevelMapping(createdBy, createApiKeyRequest.OrganisationId, constants.ORGANISATION)
	if err != nil {
		return apiKey, err
	}

	if !userProjectPrivilege.IsWrite {
		return apiKey, fmt.Errorf(string(constants.Unauthorised))
	}

	apiKeyString, err := utilities.GenerateApiKey(64)
	if err != nil {
		return apiKey, err
	}

	hashApiKey := utilities.GenerateSha512Hash(apiKeyString, constants.ApiKey)
	apiKey = createApiKeyRequest.CreateApiKeyEntity(createdBy, hashApiKey)
	err = aks.ApiKeyRepository.CreateApiKey(&apiKey)

	apiKey.ApiKey = apiKeyString
	return apiKey, err
}

func (aks *ApiKeyService) DeleteApiKey(apiKeyId string, organisationId string, deletedBy string) (err error) {
	userProjectPrivilege, err := aks.UserMappingRepository.GetUserLevelMapping(deletedBy, organisationId, constants.ORGANISATION)
	if err != nil {
		return err
	}

	if !userProjectPrivilege.IsDelete {
		return fmt.Errorf(string(constants.Unauthorised))
	}

	err = aks.ApiKeyRepository.DeleteApiKey(apiKeyId, deletedBy)

	return err
}
