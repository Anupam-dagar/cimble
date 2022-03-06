package models

import "time"

type UserPassword struct {
	UserId       string    `gorm:"primaryKey;size:255;not null"`
	PasswordHash []byte    `gorm:"size:64;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	CreatedBy    string    `gorm:"size:255;not null"`
	UpdatedBy    string    `gorm:"size:255;not null"`
}
