package repo

import (
	"dotamaster/models"
)

type vpMatch struct {
	base
}

var VpMatch vpMatch

func (v vpMatch) GetIdsExistsIn(matchIds []string) ([]string, error) {
	return []string{}, nil
}

func (v vpMatch) Create(match *models.VpMatch) error {
	return v.create(match)
}
