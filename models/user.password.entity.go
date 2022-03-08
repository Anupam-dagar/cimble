package models

type UserPassword struct {
	UserId       string `gorm:"primaryKey;size:255;not null"`
	PasswordHash []byte `gorm:"size:64;not null"`
	BaseEntity
}
