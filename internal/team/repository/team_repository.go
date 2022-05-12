package repository

import (
	"github.com/alimikegami/compnouron/internal/team/entity"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func CreateNewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (cr *TeamRepository) CreateTeam(team entity.Team) error {
	result := cr.db.Create(&team)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tr *TeamRepository) GetTeamsByUserID(ID uint) ([]entity.Team, error) {
	var teams []entity.Team
	result := tr.db.Where("user_id = ?", ID).Find(&teams)

	if result.Error != nil {
		return teams, result.Error
	}

	return teams, nil
}
