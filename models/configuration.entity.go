package models

type Configuration struct {
	ID        string `gorm:"primaryKey;size:255;not null"`
	Name      string `gorm:"size:255;not null"`
	Info      string `gorm:"size:255;not null"`
	ProjectId string `gorm:"size:255;not null"`
	BaseEntity
}
