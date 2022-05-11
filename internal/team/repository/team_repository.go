package repository

import (
	"errors"

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

func (cr *TeamRepository) UpdateTeam(team entity.Team) error {
	result := cr.db.Model(&team).Where("id = ?", team.ID).Updates(team)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no affected rows")
	}

	return nil
}
