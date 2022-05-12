package repository

import (
	"errors"
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

func (cr *TeamRepository) CreateTeam(team entity.Team) (entity.Team, error) {
	result := cr.db.Create(&team)
	if result.Error != nil {
		return team, result.Error
	}

	return team, nil
}

func (cr *TeamRepository) AddTeamMember(userID uint, teamID uint, isLeader uint) error {
	result := cr.db.Create(&entity.TeamMember{
		TeamID:   teamID,
		UserID:   userID,
		IsLeader: isLeader,
	})

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

func (tr *TeamRepository) GetTeamsByUserID(ID uint) ([]entity.Team, error) {
	var teams []entity.Team
	result := tr.db.Joins("JOIN team_members ON team_members.team_id = teams.id").Where("team_members.user_id = ?", ID).Find(&teams)

	if result.Error != nil {
		return []entity.Team{}, result.Error
	}

	return teams, nil
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
	fmt.Println(teamMembers)
	if result.Error != nil {
		return []entity.TeamMember{}, result.Error
	}

	return teamMembers, nil
}
