package models

type OrganisationModel struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CreatedAt     string `json:"createdAt"`
	ProjectsCount string `json:"projectsCount"`
	Pagination
}
