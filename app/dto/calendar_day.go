package dto

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/util"
)

type CalendarDay struct {
	Id      uint       `json:"id"`
	Type    string     `json:"type"`
	Date    string     `json:"date"`
	Workday []*Workday `json:"workday"`
}

func (receiver CalendarDay) Make(cd dao.CalendarDay) CalendarDay {
	return CalendarDay{
		Id:      cd.ID,
		Type:    cd.CalendarDayType.Name,
		Date:    util.FormatDate(cd.Date),
		Workday: MakeWorkdayList(cd.Workdays),
	}
}
