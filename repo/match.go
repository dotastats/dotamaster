package repo

import (
	"dotamaster/infra"
	"dotamaster/models"
)

type vpMatch struct {
	base
}

var VpMatch vpMatch

func (v vpMatch) GetIdsExistsIn(matchIds []int) ([]int, error) {
	if len(matchIds) == 0 {
		return []int{}, nil
	}
	var idExists []int
	err := infra.PostgreSql.Model(&models.VpMatch{}).
		Where("match_id IN (?)", matchIds).
		Pluck("match_id", &idExists).
		Error
	return idExists, err
}

func (v vpMatch) Create(match *models.VpMatch) error {
	return v.create(match)
}
