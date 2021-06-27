package dao

import (
	"gorm.io/gorm"
	"umarutv/common/models"
	"umarutv/database"
)

func CreateUserFromFullUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FullUserByName(name string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("name", name).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserByNameExists(name string) bool {
	return database.DB.Table(models.UserTableName).Select("id").Where("name", name).First(&struct{ ID int }{}).Error != gorm.ErrRecordNotFound
}
