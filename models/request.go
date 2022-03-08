package models

import "github.com/segmentio/ksuid"

type OrganisationCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (ocr OrganisationCreateRequest) CreateOrganisationEntity(createdBy string) Organisation {
	return Organisation{
		ID:        ksuid.New().String(),
		Name:      ocr.Name,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
}

type ProjectCreateRequest struct {
	OrganisationId string `json:"organisationId" binding:"required"`
	Name           string `json:"name" binding:"required"`
}

func (pcr ProjectCreateRequest) CreateProjectEntity(createdBy string) Project {
	return Project{
		ID:             ksuid.New().String(),
		OrganisationId: pcr.OrganisationId,
		Name:           pcr.Name,
		BaseEntity: BaseEntity{
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	}
}
