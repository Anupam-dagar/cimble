package models

type Project struct {
	ID             string `gorm:"primaryKey;size:255;not null"`
	OrganisationId string `gorm:"size:255;not null"`
	Name           string `gorm:"size:255;not null"`
	BaseEntity
}
