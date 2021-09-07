package main

import (
	"database/sql"
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/repository"
	"time"
)

var db = repository.DB()

func loopDaysInMonth(target time.Time, f func(t time.Time)) {
	year, month, _ := target.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
		f(d)
	}
}

func loopDaysInYear(target time.Time, f func(t time.Time)) {
	year, _, _ := target.Date()
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	for d := start; d.Year() == start.Year(); d = d.AddDate(0, 0, 1) {
		f(d)
	}
}

func insertCalendarDays() {
	db.Create(&dao.CalendarDayType{Name: "workday"})
	db.Create(&dao.CalendarDayType{Name: "holiday"})

	var workdayType, holidayType dao.CalendarDayType
	db.First(&workdayType, "name=?", "workday")
	db.First(&holidayType, "name=?", "holiday")

	createDay := func(t time.Time) {
		var dayType = workdayType
		if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
			dayType = holidayType
		}

		db.Create(&dao.CalendarDay{
			Date:            t,
			CalendarDayType: dayType,
		})
	}

	// loopDaysInMonth(time.Now(), createDay)
	loopDaysInYear(time.Now(), createDay)
}

func insertUsers() {
	db.Create(&dao.UserType{Name: "employee"})
	db.Create(&dao.UserType{Name: "boss"})

	var employeeType dao.UserType
	db.First(&employeeType, "name = ?", "employee")

	var onecareTeam dao.Team
	db.First(&onecareTeam, "name = ?", "onecare")

	db.Create(&dao.User{
		Model:           gorm.Model{},
		Name:            "Bela",
		Email:           "bela@bela.com",
		FirstDay:        time.Date(2012, time.April, 12, 0, 0, 0, 0, time.UTC),
		DefaultDayStart: "9:00",
		DefaultDayEnd:   "17:00",
		WorkHours:       8,
		UserType:        employeeType,
		UserStatus:      "active",
		Password:        "12345678",
		Teams:           []*dao.Team{&onecareTeam},
	})
	db.Create(&dao.User{
		Model:           gorm.Model{},
		Name:            "Janos",
		Email:           "janos@janos.com",
		FirstDay:        time.Date(1988, time.April, 12, 0, 0, 0, 0, time.UTC),
		LastDay:         sql.NullTime{},
		DefaultDayStart: "9:00",
		DefaultDayEnd:   "17:00",
		WorkHours:       8,
		UserType:        employeeType,
		UserStatus:      "active",
		Password:        "12345678",
		Teams:           []*dao.Team{&onecareTeam},
	})
}

func insertWorkdays() {
	db.Create(&dao.WorkdayType{Name: "workday"})
	db.Create(&dao.WorkdayType{Name: "vacation"})
	db.Create(&dao.WorkdayType{Name: "sick-leave"})

	var workdayType, vacationType, sickLeaveType dao.WorkdayType
	db.First(&workdayType, "name=?", "workday")
	db.First(&vacationType, "name=?", "vacation")
	db.First(&sickLeaveType, "name=?", "sick-leave")

	var calendarDays []dao.CalendarDay
	db.Limit(10).Find(&calendarDays)

	var user dao.User
	db.First(&user)

	db.Create(&dao.Workday{
		CalendarDay: calendarDays[0],
		WorkdayType: workdayType,
		Summary:     "Dolgoztam sokat",
		User:        user,
		Hours:       8,
	})
	db.Create(&dao.Workday{
		CalendarDay: calendarDays[1],
		WorkdayType: vacationType,
		Summary:     "Szabiztam",
		User:        user,
		Hours:       0,
	})
	db.Create(&dao.Workday{
		CalendarDay: calendarDays[2],
		WorkdayType: sickLeaveType,
		Summary:     "Haltam meg",
		User:        user,
		Hours:       0,
	})
}

func insertTeams() {
	db.Create(&dao.Team{Name: "onecare"})
	db.Create(&dao.Team{Name: "elmu"})
}

func InsertMockData() {
	insertTeams()
	insertUsers()
	insertCalendarDays()
	insertWorkdays()
}
