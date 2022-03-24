package models

type Project struct {
	ID             string        `gorm:"primaryKey;size:255;not null" json:"id"`
	OrganisationId string        `gorm:"size:255;not null" json:"organisationId"`
	Name           string        `gorm:"size:255;not null" json:"name"`
	Organisation   *Organisation `gorm:"foreignKey:ID;references:OrganisationId"`
	BaseEntity
}
