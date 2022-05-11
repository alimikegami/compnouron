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

func (cr *TeamRepository) AddTeamMember(userID uint, teamID uint) error {
	result := cr.db.Create(&entity.TeamMember{
		TeamID: teamID,
		UserID: userID,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
