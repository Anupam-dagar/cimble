package models

type ProjectModel struct {
	ID                  string `json:"id"`
	OrganisationId      string `json:"organisationId"`
	Name                string `json:"name"`
	ConfigurationsCount int    `json:"configurationsCount"`
	BaseEntity
}
