package model

import (
	"github.com/pkg/errors"
)

func GetTeamsList() ([]TeamLists, error) {
	teams := make([]TeamLists, 0)
	if err := Conn.Table("team_lists").Find(&teams).Error; err != nil {
		return nil, errors.New("战队列表获取失败")
	}
	return teams, nil
}
func GetTeamDetails(name string) (TeamLists, error) {
	var team TeamLists
	if err := Conn.Table("team_lists").Where("team_name", name).Find(&team).Error; err != nil {
		return TeamLists{}, errors.New("战队列表获取失败")
	}
	return team, nil
}
func GetTeamPlayers(title string) ([]Players, error) {
	players := make([]Players, 0)
	if err := Conn.Table("players").Where("team_name", title).Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}
func GetPlayers() ([]Players, error) {
	players := make([]Players, 0)
	if err := Conn.Table("players").Find(&players).Error; err != nil {
		return nil, errors.New("成员列表获取失败")
	}
	return players, nil
}
