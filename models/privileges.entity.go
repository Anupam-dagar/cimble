package models

type Privilege struct {
	Id       int32  `gorm:"primaryKey;autoIncrement;not null"`
	Name     string `gorm:"size:20;not null"`
	IsRead   int    `gorm:"column:tinyint;default:true;not null"`
	IsWrite  int    `gorm:"column:tinyint;default:false;not null"`
	IsUpdate int    `gorm:"column:tinyint;default:false;not null"`
	IsDelete int    `gorm:"column:tinyint;default:false;not null"`
}
