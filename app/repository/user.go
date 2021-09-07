package repository

import (
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dao"
)

type UType string

const (
	Employee UType = "employee"
	Boss     UType = "boss"
)

func FindAllUsers(users *[]dao.User) {
	db.Preload("UserType").Preload("Teams").Find(&users)
}

// FindUserById todo: szetvalasztani a preloaded es nem preloaded kerest
func FindUserById(user *dao.User, id uint) error {
	return db.Preload("UserType").First(&user, id).Error
}

func FindUserType(userType *dao.UserType, uType UType) {
	db.First(&userType, "name = ?", uType)
}

func CreateUser(user *dao.User) *gorm.DB {
	return db.Create(&user)
}

func DeleteUser(id uint) error {
	var user dao.User

	err := FindUserById(&user, id)
	if err != nil {
		return err
	}

	err = db.Model(&user).Association("Teams").Clear()
	if err != nil {
		return err
	}

	return db.Delete(&dao.User{}, id).Error
}

func SaveUser(user *dao.User) error {
	return db.Save(user).Error
}
