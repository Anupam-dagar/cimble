package models

type User struct {
	ID           string       `gorm:"primaryKey;size:255;not null"`
	FirstName    string       `gorm:"size:255;not null"`
	LastName     string       `gorm:"size:255;not null"`
	Email        string       `gorm:"size:255;not null;unique"`
	UserPassword UserPassword `gorm:"foreignKey:UserId"`
	BaseEntity
}
