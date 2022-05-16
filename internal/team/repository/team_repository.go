package repository

import (
	"errors"

	"github.com/alimikegami/compnouron/internal/team/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeamRepository interface {
	CreateTeam(team entity.Team) (entity.Team, error)
	AddTeamMember(userID uint, teamID uint, isLeader uint) error
	UpdateTeam(team entity.Team) error
	DeleteTeam(id uint) error
	GetTeamsByUserID(ID uint) ([]entity.Team, error)
	GetTeamByID(teamID uint) (entity.Team, error)
	GetTeamMembersByID(teamID uint) ([]entity.TeamMember, error)
}

type TeamRepositoryImpl struct {
	db *gorm.DB
}

func CreateNewTeamRepository(db *gorm.DB) TeamRepository {
	return &TeamRepositoryImpl{db: db}
}

func (cr *TeamRepositoryImpl) CreateTeam(team entity.Team) (entity.Team, error) {
	result := cr.db.Create(&team)
	if result.Error != nil {
		return team, result.Error
	}

	return team, nil
}

func (cr *TeamRepositoryImpl) AddTeamMember(userID uint, teamID uint, isLeader uint) error {
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

func (cr *TeamRepositoryImpl) UpdateTeam(team entity.Team) error {
	result := cr.db.Model(&team).Updates(team)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no affected rows")
	}

	return nil
}

func (cr *TeamRepositoryImpl) DeleteTeam(id uint) error {
	result := cr.db.Delete(&entity.Team{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (tr *TeamRepositoryImpl) GetTeamsByUserID(ID uint) ([]entity.Team, error) {
	var teams []entity.Team
	result := tr.db.Debug().Joins("JOIN team_members ON team_members.team_id = teams.id").Where("team_members.user_id = ?", ID).Find(&teams)

	if result.Error != nil {
		return []entity.Team{}, result.Error
	}

	return teams, nil
}

func (tr *TeamRepositoryImpl) GetTeamByID(teamID uint) (entity.Team, error) {
	var team entity.Team
	result := tr.db.First(&team, teamID)
	if result.Error != nil {
		return entity.Team{}, result.Error
	}

	return team, nil
}

func (tr *TeamRepositoryImpl) GetTeamMembersByID(teamID uint) ([]entity.TeamMember, error) {
	var teamMembers []entity.TeamMember

	result := tr.db.Debug().Preload(clause.Associations).Where("team_id = ?", teamID).Find(&teamMembers)
	if result.Error != nil {
		return []entity.TeamMember{}, result.Error
	}

	return teamMembers, nil
}
