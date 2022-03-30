package models

type Configuration struct {
	ID        string `gorm:"primaryKey;size:255;not null" json:"id,omitempty"`
	Name      string `gorm:"size:255;not null" json:"name,omitempty"`
	Info      string `gorm:"size:255;not null" json:"info,omitempty"`
	ProjectId string `gorm:"size:255;not null" json:"projectId,omitempty"`
	BaseEntity
}

func (c Configuration) CreateConfigurationArchiveEntity(deletedBy string) ConfigurationArchive {
	return ConfigurationArchive{
		Configuration: c,
		DeletedBy:     deletedBy,
	}
}
