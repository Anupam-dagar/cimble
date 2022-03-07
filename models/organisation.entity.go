package models

import "time"

type Organisation struct {
	ID        string    `gorm:"primaryKey;size:255;not null"`
	Name      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedBy string    `gorm:"size:255;not null"`
	UpdatedBy string    `gorm:"size:255;not null"`
}
