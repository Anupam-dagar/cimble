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

type OrganisationUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (our OrganisationUpdateRequest) CreateUpdateOrgnisationEntity(updatedBy string) Organisation {
	return Organisation{
		Name:      our.Name,
		UpdatedBy: updatedBy,
	}
}

type ProjectUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (pur ProjectUpdateRequest) CreateUpdateProjectEntity(updatedBy string) Project {
	return Project{
		Name: pur.Name,
		BaseEntity: BaseEntity{
			UpdatedBy: updatedBy,
		},
	}
}

type ConfigurationCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info" binding:"required"`
}

func (ccr ConfigurationCreateRequest) CreateConfigurationEntity(createdBy string, projectId string) Configuration {
	return Configuration{
		ID:        ksuid.New().String(),
		Name:      ccr.Name,
		Info:      ccr.Info,
		ProjectId: projectId,
		BaseEntity: BaseEntity{
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	}
}

type ConfigurationUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info" binding:"required"`
}

func (cur ConfigurationUpdateRequest) CreateUpdateConfigurationEntity(updatedBy string) Configuration {
	return Configuration{
		Name: cur.Name,
		Info: cur.Info,
		BaseEntity: BaseEntity{
			UpdatedBy: updatedBy,
		},
	}
}

type PaginationQueryParams struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}
