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
	GetApiKeys(string, string, int64, int64) (models.ApiKeysResponse, error)
	IsValidApiKey(string, string) (bool, error)
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

func (aks *ApiKeyService) GetApiKeys(organisationId string, userId string, offset int64, limit int64) (apiKeys models.ApiKeysResponse, err error) {
	userProjectPrivilege, err := aks.UserMappingRepository.GetUserLevelMapping(userId, organisationId, constants.ORGANISATION)
	if err != nil {
		return apiKeys, err
	}

	if !userProjectPrivilege.IsRead {
		return apiKeys, fmt.Errorf(string(constants.Unauthorised))
	}

	apiKeysData, count, err := aks.ApiKeyRepository.GetApiKeys(organisationId, offset, limit)
	pagination := utilities.GeneratePage(count, offset, limit)

	apiKeys = models.ApiKeysResponse{
		ApiKeys: apiKeysData,
		Page:    pagination,
	}
	return apiKeys, err
}

func (aks *ApiKeyService) IsValidApiKey(organisationId string, apiKey string) (isValid bool, err error) {
	hashApiKey := utilities.GenerateSha512Hash(apiKey, constants.ApiKey)
	isValid, err = aks.ApiKeyRepository.ValidateApiKey(organisationId, hashApiKey)

	return isValid, err
}
