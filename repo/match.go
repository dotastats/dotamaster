package repo

import (
	"dotamaster/infra"
	"dotamaster/models"
	"dotamaster/utils/uerror"

	"github.com/jinzhu/gorm"
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

func (v vpMatch) GetByMatchId(matchId int) (*models.VpMatch, error) {
	object := models.VpMatch{}
	err := infra.PostgreSql.
		Where("match_id = ?", matchId).
		Find(&object).
		Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &object, err
}

func (v vpMatch) Create(match *models.VpMatch) error {
	return v.create(match)
}

func (v vpMatch) Upsert(newMatch *models.VpMatch) error {
	match, err := v.GetByMatchId(newMatch.MatchID)
	if err != nil {
		return uerror.StackTrace(err)
	}

	if match == nil {
		return v.create(match)
	}

	return v.save(newMatch)
}
