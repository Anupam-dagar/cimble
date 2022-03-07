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
