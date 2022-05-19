package repository

import (
	"errors"

	"github.com/alimikegami/compnouron/internal/team/entity"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreateTeam(team entity.Team) (entity.Team, error)
	AddTeamMember(userID uint, teamID uint, isLeader uint) error
	UpdateTeam(team entity.Team) error
	DeleteTeam(id uint) error
	GetTeamsByUserID(ID uint) ([]entity.Team, error)
	GetTeamByID(teamID uint) (entity.Team, error)
	GetTeamLeader(teamID uint) (uint, error)
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

func (cr *TeamRepositoryImpl) GetTeamLeader(teamID uint) (uint, error) {
	var leader entity.TeamMember
	result := cr.db.First(&leader, "team_id = ? AND is_leader = 1", teamID)

	if result.Error != nil {
		return 0, result.Error
	}

	return leader.UserID, nil
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

	result := tr.db.Debug().Preload("TeamMembers.User").Preload("TeamMembers").Find(&team, teamID)
	if result.Error != nil {
		return entity.Team{}, result.Error
	}

	return team, nil
}
