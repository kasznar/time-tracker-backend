package services

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/repository"
	"time"
)

func ListAllCalendarDays() []dto.CalendarDay {
	var days []dao.CalendarDay
	var response []dto.CalendarDay

	repository.FindAllCalendarDays(&days)

	for _, day := range days {
		var dayDto = dto.CalendarDay{}.Make(day)
		response = append(response, dayDto)
	}

	return response
}

func ListDaysInMonth(year int, month time.Month) []dto.CalendarDay {
	var days []dao.CalendarDay
	var response []dto.CalendarDay

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)

	repository.FindCalendarDaysBetweenDates(&days, start, end)

	for _, day := range days {
		var dayDto = dto.CalendarDay{}.Make(day)
		response = append(response, dayDto)
	}

	return response
}
