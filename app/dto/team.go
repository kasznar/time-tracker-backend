package dto

import "kasznar.hu/hello/app/dao"

type TeamDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (r TeamDto) Make(team *dao.Team) TeamDto {
	return TeamDto{
		Id:   team.ID,
		Name: team.Name,
	}
}

type TeamListDto = []*TeamDto

// (r TeamListDto)
func MakeTeamList(teams []*dao.Team) (teamList TeamListDto) {
	for _, t := range teams {
		dto := TeamDto{}.Make(t)
		teamList = append(teamList, &dto)
	}

	return
}

type CreateTeam struct {
	Name string `json:"name"`
}

type DeleteTeam struct {
	Id uint `json:"id"`
}
