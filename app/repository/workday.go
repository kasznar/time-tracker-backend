package repository

import (
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dao"
	"time"
)

type WDType string

const (
	Workday   WDType = "workday"
	Vacation  WDType = "vacation"
	SickLeave WDType = "sick-leave"
)

func FindAllWorkdays(workdays *[]dao.Workday) {
	db.Preload("WorkdayType").Preload("User").Preload("CalendarDay").Find(&workdays)
}

func FindWorkdayById(workday *dao.Workday, id uint) error {
	return db.First(workday, id).Error
}

func CreateWorkDay(workday *dao.Workday) *gorm.DB {
	return db.Create(&workday)
}

func UpdateWorkDay(workday *dao.Workday) error {
	return db.Save(workday).Error
}

func FindWorkdayType(workdayType *dao.WorkdayType, wdType WDType) {
	db.First(&workdayType, "name = ?", wdType)
}

func FindWorkdaysByUser(workdays *[]dao.Workday, id uint) {
	db.Preload("WorkdayType").Preload("CalendarDay").Where("user_id = ?", id).Find(&workdays)
}

func FindWorkdaysByUserAndDate(days *[]dao.CalendarDay, userId uint, startDate, endDate time.Time) {
	db.
		Preload("Workdays.WorkdayType").
		Preload("Workdays", "user_id = ?", userId).
		Preload("CalendarDayType").
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Find(days)
}

func FindWorkdaysByTeamAndDate(teamWorkdays *[]dao.APIWorkdaysInTeamByDate, teamId uint, startDate, endDate time.Time) {
	// language=SQL
	query := `select calendar_days.date, users.id, users.name as name, w.id as workday_id, w.summary, wt.name as workday_type
				from calendar_days
    			cross join users
    			left join workdays w on calendar_days.id = w.calendar_day_id and users.id = w.user_id
				left join workday_types wt on w.workday_type_id = wt.id
				left join user_teams ut on users.id = ut.user_id
				where
      			ut.team_id = ? and
      			calendar_days.date between ? and ?
      			order by calendar_days.date`

	db.Raw(query, teamId, startDate, endDate).Scan(&teamWorkdays)
}
