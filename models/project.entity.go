package models

type Project struct {
	ID             string           `gorm:"primaryKey;size:255;not null" json:"id"`
	OrganisationId string           `gorm:"size:255;not null" json:"organisationId"`
	Name           string           `gorm:"size:255;not null" json:"name"`
	Organisation   *Organisation    `gorm:"foreignKey:ID;references:OrganisationId"`
	Configuration  []*Configuration `gorm:"foreignKey:ProjectId;references:ID" json:"configurations"`
	BaseEntity
}

func (p Project) CreateProjectArchiveEntity(deletedBy string) ProjectArchive {
	return ProjectArchive{
		Project: p,
		DeletedBaseEntity: DeletedBaseEntity{
			DeletedBy: deletedBy,
		},
	}
}
