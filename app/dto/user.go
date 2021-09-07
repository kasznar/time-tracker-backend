package dto

import (
	"database/sql"
	"kasznar.hu/hello/app/dao"
)

type UserDto struct {
	Id              uint           `json:"id"`
	Email           string         `json:"email"`
	Name            string         `json:"name"`
	FirstDay        string         `json:"firstDay"`
	DefaultDayStart string         `json:"defaultDayStart"`
	DefaultDayEnd   string         `json:"defaultDayEnd"`
	WorkHours       int            `json:"workHours"`
	UserType        string         `json:"userType"`
	Status          string         `json:"status"`
	LastDay         sql.NullString `json:"lastDay"`
	Teams           TeamListDto    `json:"teams"`
}

func (receiver UserDto) Make(user dao.User) UserDto {
	return UserDto{
		Id:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		FirstDay:        "",
		DefaultDayStart: "",
		DefaultDayEnd:   "",
		WorkHours:       user.WorkHours,
		UserType:        user.UserType.Name,
		Status:          "",
		Teams:           MakeTeamList(user.Teams),
	}
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"`
	FirstDay string `json:"firstDay"`
	UserType string `json:"userType" binding:"required"`
	Status   string `json:"status"`
}
