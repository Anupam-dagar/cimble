package repositories

import (
	"cimble/constants"
	"cimble/models"
	"cimble/utilities"

	"gorm.io/gorm"
)

type UserMappingRepositoryInterface interface {
	CreateUserMapping(*models.UserMapping, *gorm.DB) error
	GetUserLevelMapping(string, string, constants.PrivilegeLevel) (models.PrivilegeModel, error)
}

type UserMappingRepository struct {
	db *gorm.DB
}

func NewUserMappingRepository() UserMappingRepositoryInterface {
	upr := new(UserMappingRepository)
	upr.db = utilities.GetDatabase()
	return upr
}

func (upr *UserMappingRepository) CreateUserMapping(userMapping *models.UserMapping, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = upr.db
	}

	tx = tx.Table("user_mappings")
	err = tx.Create(userMapping).Error

	return err
}

func (upr *UserMappingRepository) GetUserLevelMapping(
	userId string,
	levelId string,
	levelFor constants.PrivilegeLevel,
) (privilegeModel models.PrivilegeModel, err error) {
	db := upr.db
	var privilege models.Privilege

	db = db.Table("user_mappings")
	db.Select("privileges.*")
	db.Joins("inner join privileges on privileges.name = user_mappings.privelege")
	db.Where("user_mappings.user_id = ?", userId)
	db.Where("user_mappings.level_for = ?", levelFor)
	db.Where("user_mappings.level_id = ?", levelId)

	err = db.Find(&privilege).Error
	privilegeModel = privilege.ToPrivilegeModel()

	return privilegeModel, err
}
