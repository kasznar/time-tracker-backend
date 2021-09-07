package dao

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_teams;"`
}
