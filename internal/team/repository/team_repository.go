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

func (cr *TeamRepository) DeleteTeam(id uint) error {
	result := cr.db.Delete(&entity.Team{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no rows affected")
	}

	return nil
}
