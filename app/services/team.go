package services

import (
	"kasznar.hu/hello/app/dao"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/repository"
)

func ListAllTeams() dto.TeamListDto {
	var teams []*dao.Team
	repository.FindAllTeams(&teams)

	return dto.MakeTeamList(teams)
}

func CreateTeam(request dto.CreateTeam) (id *uint, err error) {
	team := dao.Team{Name: request.Name}
	result := repository.CreateTeam(&team)

	if result.Error != nil {
		return nil, result.Error
	}

	return &team.ID, nil
}

func DeleteTeam(id uint) error {
	var team dao.Team
	err := repository.FindTeamById(&team, id)
	if err != nil {
		return err
	}
	return repository.DeleteTeam(&team)
}

func ListUsersInTeam(id uint) (response []dto.UserDto, err error) {
	var team dao.Team
	err = repository.FindTeamById(&team, id)

	for _, user := range team.Users {
		var userDto = dto.UserDto{}.Make(*user)
		response = append(response, userDto)
	}

	return
}

func AddUserToTeam(teamId uint, userId uint) error {
	var team dao.Team
	err := repository.FindTeamById(&team, teamId)
	if err != nil {
		return err
	}

	var user dao.User
	err = repository.FindUserById(&user, userId)
	if err != nil {
		return err
	}

	user.Teams = append(user.Teams, &team)
	err = repository.SaveUser(&user)
	if err != nil {
		return err
	}

	return nil
}
