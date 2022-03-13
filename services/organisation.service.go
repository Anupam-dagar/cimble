package services

import (
	"cimble/constants"
	"cimble/models"
	"cimble/repositories"
	"fmt"
)

type OrganisationServiceInterface interface {
	CreateOrganisation(models.OrganisationCreateRequest, string) (models.Organisation, error)
	UpdateOrganisation(models.OrganisationUpdateRequest, string, string) (models.Organisation, error)
	GetOrganisations(string) ([]models.Organisation, error)
}

type OrganisationService struct {
	OrganisationRepository repositories.OrganisationRepositoryInterface
	UserMappingRepository  repositories.UserMappingRepositoryInterface
}

func NewOrganisationService() OrganisationServiceInterface {
	os := new(OrganisationService)
	os.OrganisationRepository = repositories.NewOrganisationRepository()
	os.UserMappingRepository = repositories.NewUserMappingRepository()
	return os
}

func (os *OrganisationService) CreateOrganisation(
	organisationPayload models.OrganisationCreateRequest,
	createdBy string,
) (organisation models.Organisation, err error) {
	organisation = organisationPayload.CreateOrganisationEntity(createdBy)
	userMapping := models.UserMapping{
		UserId:    createdBy,
		LevelFor:  string(constants.ORGANISATION),
		LevelId:   organisation.ID,
		Privelege: string(constants.OWNER),
		BaseEntity: models.BaseEntity{
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	}
	err = os.OrganisationRepository.CreateOrganisation(&organisation, &userMapping)
	if err != nil {
		fmt.Printf("error creating organisation: %v", err)
		return organisation, err
	}

	return organisation, err
}

func (os *OrganisationService) UpdateOrganisation(
	organisationPayload models.OrganisationUpdateRequest,
	organisationId,
	updatedBy string,
) (organisation models.Organisation, err error) {
	organisation = organisationPayload.CreateUpdateOrgnisationEntity(updatedBy)
	userOrganisationPrivilege, err := os.UserMappingRepository.GetUserOrganisationMapping(updatedBy, organisationId)

	if err != nil {
		return organisation, err
	}

	if !userOrganisationPrivilege.IsUpdate {
		return organisation, fmt.Errorf(string(constants.Unauthorised))
	}

	err = os.OrganisationRepository.UpdateOrganisationById(&organisation, organisationId)
	if err != nil {
		fmt.Printf("error updating organisation %s by %s: %v", organisationId, updatedBy, err)
		return organisation, err
	}

	return organisation, err
}

func (os *OrganisationService) GetOrganisations(userId string) (organisations []models.Organisation, err error) {
	organisations, err = os.OrganisationRepository.GetOrganisations(userId)
	if err != nil {
		return organisations, err
	}

	return organisations, err
}
