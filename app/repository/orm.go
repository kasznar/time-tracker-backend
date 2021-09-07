package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dao"
	"os"
)

var db *gorm.DB

func Connect() {
	user := os.Getenv("TIME_TRACKER_DB_USER")
	password := os.Getenv("TIME_TRACKER_DB_PASSWORD")
	//db, err := gorm.Open(mysql.Open("GO_DB"), &gorm.Config{})
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/GO_DB?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func MigrateOrm() {
	err := db.AutoMigrate(
		&dao.UserType{},
		&dao.User{},
		&dao.CalendarDayType{},
		&dao.CalendarDay{},
		&dao.Workday{},
		&dao.WorkdayType{},
		&dao.Team{},
	)

	if err != nil {
		panic("failed to migrate")
	}
}

func DB() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}
