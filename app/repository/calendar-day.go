package repository

import (
	"kasznar.hu/hello/app/dao"
	"time"
)

type CDayType string

const (
	WorkDay CDayType = "workday"
	Holiday CDayType = "holiday"
)

func FindAllCalendarDays(calendarDays *[]dao.CalendarDay) {
	db.Preload("CalendarDayType").Find(&calendarDays)
}

func FindCalendarDayByID(day *dao.CalendarDay, id int) error {
	return db.First(&day, id).Error
}

func FindCalendarDaysBetweenDates(days *[]dao.CalendarDay, startDate, endDate time.Time) {
	db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&days)
}
