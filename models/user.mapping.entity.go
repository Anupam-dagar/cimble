package models

type UserMapping struct {
	UserId    string `gorm:"primaryKey;size:255;not null"`
	LevelFor  string `gorm:"default:project;not null"`
	LevelId   string `gorm:"primaryKey;size:255;not null"`
	Privelege string `gorm:"size:20;not null"`
	BaseEntity
}
