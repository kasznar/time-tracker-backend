package dao

import (
	"gorm.io/gorm"
	"time"
)

type CalendarDayType struct {
	gorm.Model
	Name string
}

type CalendarDay struct {
	gorm.Model
	Date              time.Time
	CalendarDayTypeID uint
	CalendarDayType   CalendarDayType
	Workdays          []*Workday
}
