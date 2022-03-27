package models

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	CreatedBy string     `gorm:"size:255;not null" json:"createdBy,omitempty"`
	UpdatedBy string     `gorm:"size:255;not null" json:"updatedBy,omitempty"`
}
