package services

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/repository"
	"kasznar.hu/hello/app/util"
)

func ListAllUsers() (response []dto.UserDto) {
	var users []dao.User

	repository.FindAllUsers(&users)

	for _, user := range users {
		var userDto = dto.UserDto{}.Make(user)
		response = append(response, userDto)
	}

	return
}

func FindUserById(id uint) (user dto.UserDto, err error) {
	var userDao dao.User
	err = repository.FindUserById(&userDao, id)
	user = dto.UserDto{}.Make(userDao)

	return
}

func CreateUser(request dto.CreateUserRequest) (id *uint, err error) {
	var employeeType dao.UserType
	repository.FindUserType(&employeeType, repository.Employee)

	var user = dao.User{
		Name:            request.Name,
		Email:           request.Email,
		FirstDay:        util.DateFromMidnight(2012, 12, 12),
		DefaultDayStart: "9:00:00",
		DefaultDayEnd:   "17:00:00",
		WorkHours:       8,
		UserType:        employeeType,
		UserStatus:      "",
		Password:        request.Password,
	}

	result := repository.CreateUser(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user.ID, nil
}

func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}
