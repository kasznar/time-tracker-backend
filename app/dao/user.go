package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type UserType struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Name            string
	Email           string
	FirstDay        time.Time
	LastDay         sql.NullTime
	DefaultDayStart string
	DefaultDayEnd   string
	WorkHours       int
	UserTypeID      uint
	UserType        UserType
	UserStatus      string
	Password        string
	Teams           []*Team `gorm:"many2many:user_teams;"`
}
