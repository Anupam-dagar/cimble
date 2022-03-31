package models

import "time"

type Organisation struct {
	ID        string     `gorm:"primaryKey;size:255;not null" json:"id"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	CreatedBy string     `gorm:"size:255;not null" json:"createdBy"`
	UpdatedBy string     `gorm:"size:255;not null" json:"updatedBy"`
	Projects  []*Project `gorm:"foreignKey:OrganisationId;references:ID" json:"projects"`
}

func (o Organisation) CreateOrganisationArchiveEntity(deletedBy string) OrganisationArchive {
	return OrganisationArchive{
		Organisation: o,
		DeletedBaseEntity: DeletedBaseEntity{
			DeletedBy: deletedBy,
		},
	}
}
