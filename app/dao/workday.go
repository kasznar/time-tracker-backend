package dao

import (
	"gorm.io/gorm"
	"time"
)

type WorkdayType struct {
	gorm.Model
	Name string
}

type Workday struct {
	gorm.Model
	CalendarDayID uint
	CalendarDay   CalendarDay
	WorkdayTypeID uint
	WorkdayType   WorkdayType
	Summary       string
	UserID        uint
	User          User
	Hours         uint
}

type APIWorkdaysInTeamByDate struct {
	Date        time.Time `json:"date"`
	ID          uint      `json:"user_id"`
	Name        string    `json:"user_name"`
	WorkdayId   *uint     `json:"workday_id"`
	Summary     *string   `json:"summary"`
	WorkdayType *string   `json:"workday_type"`
}
