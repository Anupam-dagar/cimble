package models

import "time"

type UserOrganisationMapping struct {
	ID             string    `gorm:"primaryKey;size:255;not null"`
	UserId         string    `gorm:"size:255;not null"`
	OrganisationId string    `gorm:"size:255;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	CreatedBy      string    `gorm:"size:255;not null"`
	UpdatedBy      string    `gorm:"size:255;not null"`
}
