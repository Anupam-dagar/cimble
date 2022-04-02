package models

type ApiKey struct {
	ID             string `gorm:"primaryKey;size:255;not null" json:"id"`
	OrganisationId string `gorm:"size:255;not null" json:"organisationId"`
	KeyHash        []byte `gorm:"size:64;not null" json:"-"`
	Privileges     string `gorm:"size:255;not null" json:"privileges"`
	Revoked        int    `gorm:"default:0;not null" json:"revoked"`
	ApiKey         string `gorm:"-:all" json:"apiKey"`
	BaseEntity
}
