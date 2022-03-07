package services

import (
	"cimble/models"
	"cimble/repositories"
	"fmt"

	"github.com/segmentio/ksuid"
)

type OrganisationServiceInterface interface {
	CreateOrganisation(models.OrganisationCreateRequest, string) (models.Organisation, error)
}

type OrganisationService struct {
	OrganisationRepository repositories.OrganisationRepositoryInterface
}

func NewOrganisationService() OrganisationServiceInterface {
	os := new(OrganisationService)
	os.OrganisationRepository = repositories.NewOrganisationRepository()
	return os
}

func (os *OrganisationService) CreateOrganisation(
	organisationPayload models.OrganisationCreateRequest,
	createdBy string,
) (organisation models.Organisation, err error) {
	organisation = organisationPayload.CreateOrganisationEntity(createdBy)
	userOrganisationMapping := models.UserOrganisationMapping{
		ID:             ksuid.New().String(),
		UserId:         createdBy,
		OrganisationId: organisation.ID,
		CreatedBy:      createdBy,
		UpdatedBy:      createdBy,
	}

	err = os.OrganisationRepository.CreateOrganisation(organisation, userOrganisationMapping)
	if err != nil {
		fmt.Printf("error creating organisation: %v", err)
		return organisation, err
	}

	return organisation, err
}
