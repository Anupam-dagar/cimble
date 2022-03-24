package models

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	CreatedBy string    `gorm:"size:255;not null" json:"createdBy"`
	UpdatedBy string    `gorm:"size:255;not null" json:"updatedBy"`
}
