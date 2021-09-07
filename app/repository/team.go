package repository

import (
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dao"
)

func FindAllTeams(teams *[]*dao.Team) {
	db.Find(&teams)
}

func CreateTeam(team *dao.Team) *gorm.DB {
	return db.Create(&team)
}

func DeleteTeam(team *dao.Team) error {
	err := db.Model(team).Association("Users").Clear()
	if err != nil {
		return err
	}
	return db.Delete(team).Error
}

func FindTeamById(team *dao.Team, id uint) error {
	return db.Preload("Users").First(&team, id).Error
}
