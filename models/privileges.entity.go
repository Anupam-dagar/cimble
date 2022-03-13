package models

type Privilege struct {
	Id       int32  `gorm:"primaryKey;autoIncrement;not null"`
	Name     string `gorm:"size:20;not null"`
	IsRead   int    `gorm:"default:1;not null"`
	IsWrite  int    `gorm:"default:0;not null"`
	IsUpdate int    `gorm:"default:0;not null"`
	IsDelete int    `gorm:"default:0;not null"`
}

type PrivilegeModel struct {
	Id       int32
	Name     string
	IsRead   bool
	IsWrite  bool
	IsUpdate bool
	IsDelete bool
}

func (p Privilege) ToPrivilegeModel() PrivilegeModel {
	return PrivilegeModel{
		Id:       p.Id,
		Name:     p.Name,
		IsRead:   p.IsRead == 1,
		IsWrite:  p.IsWrite == 1,
		IsUpdate: p.IsUpdate == 1,
		IsDelete: p.IsDelete == 1,
	}
}
