package dto

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/repository"
	"kasznar.hu/hello/app/util"
)

type Workday struct {
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Date    string `json:"calendar_day"`
	UserId  uint   `json:"user_id"`
}

func (receiver Workday) Make(workday dao.Workday) Workday {
	return Workday{
		Id:      workday.ID,
		Type:    workday.WorkdayType.Name,
		Summary: workday.Summary,
		Date:    util.FormatDate(workday.CalendarDay.Date),
		UserId:  workday.User.ID,
	}
}

func MakeWorkdayList(workdaysDao []*dao.Workday) (workdays []*Workday) {
	for _, d := range workdaysDao {
		dto := Workday{}.Make(*d)
		workdays = append(workdays, &dto)
	}

	return
}

type ListByCalendarDayWorkday struct {
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
}

type ListByCalendarDay struct {
	CalendarDayId uint                      `json:"calendar_day_id"`
	Date          string                    `json:"date"`
	Type          string                    `json:"type"`
	Workday       *ListByCalendarDayWorkday `json:"workday"`
}

type CrateWorkdayRequest struct {
	Summary string            `json:"summary"`
	Type    repository.WDType `json:"type" binding:"required"`
	DayId   int               `json:"day_id" binding:"required"`
	UserId  uint              `json:"user_id" binding:"required"`
}

type UpdateWorkdayRequest struct {
	Summary   *string            `json:"summary"`
	Type      *repository.WDType `json:"type"`
	WorkdayId uint               `json:"workday_id" binding:"required"`
}

type ListWorkDaysRequest struct {
	UserId *uint `json:"user_id"`
	Year   *int  `json:"year"`
	Month  *int  `json:"month"`
}
