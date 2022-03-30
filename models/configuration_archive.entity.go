package models

import "time"

type ConfigurationArchive struct {
	Configuration
	DeletedTimestamp *time.Time `gorm:"autoCreateTime" json:"deletedAt,omitempty"`
	DeletedBy        string     `gorm:"size:255;not null" json:"deletedBy,omitempty"`
}
