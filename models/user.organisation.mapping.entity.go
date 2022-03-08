package models

type UserOrganisationMapping struct {
	ID             string `gorm:"primaryKey;size:255;not null"`
	UserId         string `gorm:"size:255;not null"`
	OrganisationId string `gorm:"size:255;not null"`
	BaseEntity
}
