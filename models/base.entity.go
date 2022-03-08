package models

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedBy string    `gorm:"size:255;not null"`
	UpdatedBy string    `gorm:"size:255;not null"`
}
