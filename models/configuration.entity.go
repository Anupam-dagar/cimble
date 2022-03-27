package models

type Configuration struct {
	ID        string `gorm:"primaryKey;size:255;not null" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Info      string `gorm:"size:255;not null" json:"info"`
	ProjectId string `gorm:"size:255;not null" json:"projectId"`
	BaseEntity
}
