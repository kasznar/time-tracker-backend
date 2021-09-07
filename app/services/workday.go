package services

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/repository"
	"kasznar.hu/hello/app/util"
	"time"
)

func ListAllWorkDays() (response []dto.Workday) {
	var days []dao.Workday

	repository.FindAllWorkdays(&days)

	for _, day := range days {
		var dayDto = dto.Workday{}.Make(day)
		response = append(response, dayDto)
	}

	return
}

func ListWorkDaysByUser(userId uint) (response []dto.Workday) {
	var days []dao.Workday

	repository.FindWorkdaysByUser(&days, userId)

	for _, day := range days {
		var dayDto = dto.Workday{}.Make(day)
		response = append(response, dayDto)
	}

	return
}

func ListWorkDaysByUserAndMonth(userId uint, year int, month time.Month) (response []dto.ListByCalendarDay) {
	var days []dao.CalendarDay

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)

	repository.FindWorkdaysByUserAndDate(&days, userId, start, end)

	for _, day := range days {
		var workday *dto.ListByCalendarDayWorkday

		if len(day.Workdays) > 0 {
			if wd := day.Workdays[0]; wd != nil {
				workday = &dto.ListByCalendarDayWorkday{
					Id:      wd.ID,
					Type:    wd.WorkdayType.Name,
					Summary: wd.Summary,
				}
			}
		}

		var dayDto = dto.ListByCalendarDay{
			CalendarDayId: day.ID,
			Date:          util.FormatDate(day.Date),
			Type:          day.CalendarDayType.Name,
			Workday:       workday,
		}
		response = append(response, dayDto)
	}

	return
}

func ListWorkDaysByTeamAndMonth(teamId uint, year int, month time.Month) (response []dao.APIWorkdaysInTeamByDate) {
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)

	repository.FindWorkdaysByTeamAndDate(&response, teamId, start, end)

	return
}

func CreateWorkday(request dto.CrateWorkdayRequest) (id *uint, err error) {
	var workdayType dao.WorkdayType
	var calendarDay dao.CalendarDay
	var user dao.User

	repository.FindWorkdayType(&workdayType, request.Type)
	repository.FindCalendarDayByID(&calendarDay, request.DayId)
	repository.FindUserById(&user, request.UserId)

	var workday = dao.Workday{
		CalendarDay: calendarDay,
		WorkdayType: workdayType,
		Summary:     request.Summary,
		User:        user,
		Hours:       8,
	}

	result := repository.CreateWorkDay(&workday)

	if result.Error != nil {
		return nil, result.Error
	}

	return &workday.ID, nil
}

func UpdateWorkday(request dto.UpdateWorkdayRequest) (id *uint, err error) {
	var workdayType dao.WorkdayType
	var workday dao.Workday

	err = repository.FindWorkdayById(&workday, request.WorkdayId)
	if err != nil {
		return nil, err
	}

	if request.Type != nil {
		repository.FindWorkdayType(&workdayType, *request.Type)
		workday.WorkdayType = workdayType
	}

	if request.Summary != nil {
		workday.Summary = *request.Summary
	}

	err = repository.UpdateWorkDay(&workday)
	if err != nil {
		return nil, err
	}

	return &workday.ID, nil
}
