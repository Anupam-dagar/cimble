package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;size:255;not null"`
	FirstName string    `gorm:"size:255;not null"`
	LastName  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedBy string    `gorm:"size:255;not null"`
	UpdatedBy string    `gorm:"size:255;not null"`
}
