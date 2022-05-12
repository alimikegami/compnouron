package repository

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/team/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (tr *TeamRepository) GetTeamByID(teamID uint) (entity.Team, error) {
	var team entity.Team
	result := tr.db.Model(entity.Team{ID: teamID}).First(&team)
	fmt.Println(result)
	if result.Error != nil {
		return entity.Team{}, result.Error
	}

	return team, nil
}

func (tr *TeamRepository) GetTeamMembersByID(teamID uint) ([]entity.TeamMember, error) {
	var teamMembers []entity.TeamMember

	result := tr.db.Preload(clause.Associations).Where("team_id = ?", teamID).Find(&teamMembers)
	if result.Error != nil {
		return []entity.TeamMember{}, result.Error
	}

	return teamMembers, nil
}
