package usecase

import (
	"errors"

	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/internal/team/entity"
	"github.com/alimikegami/compnouron/internal/team/repository"
)

type TeamUseCase interface {
	CreateTeam(userID uint, team dto.TeamRequest) error
	DeleteTeam(id uint, userID uint) error
	UpdateTeam(userID uint, team dto.TeamRequest, teamID uint) error
	GetTeamsByUserID(userID uint) ([]dto.BriefTeamResponse, error)
	GetTeamDetailsByID(teamID uint) (dto.TeamDetailsResponse, error)
}

type TeamUseCaseImpl struct {
	tr repository.TeamRepository
}

func CreateNewTeamUseCase(tr repository.TeamRepository) TeamUseCase {
	return &TeamUseCaseImpl{tr: tr}
}

func (tuc *TeamUseCaseImpl) CreateTeam(userID uint, team dto.TeamRequest) error {
	teamEntity := entity.Team{
		Name:        team.Name,
		Description: team.Description,
		Capacity:    team.Capacity,
	}

	teamEntity, err := tuc.tr.CreateTeam(teamEntity)
	if err != nil {
		return err
	}

	err = tuc.tr.AddTeamMember(userID, teamEntity.ID, 1)

	return err
}

func (tuc *TeamUseCaseImpl) DeleteTeam(id uint, userID uint) error {
	teamOwner, err := tuc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	err = tuc.tr.DeleteTeam(id)

	return err
}

func (tuc *TeamUseCaseImpl) UpdateTeam(userID uint, team dto.TeamRequest, teamID uint) error {
	teamOwner, err := tuc.tr.GetTeamLeader(teamID)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	teamEntity := entity.Team{
		ID:          teamID,
		Name:        team.Name,
		Description: team.Description,
		Capacity:    team.Capacity,
	}

	err = tuc.tr.UpdateTeam(teamEntity)
	return err
}

func (tuc *TeamUseCaseImpl) GetTeamsByUserID(userID uint) ([]dto.BriefTeamResponse, error) {
	var teamsResponse []dto.BriefTeamResponse
	result, err := tuc.tr.GetTeamsByUserID(userID)
	if err != nil {
		return teamsResponse, err
	}

	for _, teamResponse := range result {
		teamsResponse = append(teamsResponse, dto.BriefTeamResponse{
			ID:   teamResponse.ID,
			Name: teamResponse.Name,
		})
	}

	return teamsResponse, nil
}

func (tuc *TeamUseCaseImpl) GetTeamDetailsByID(teamID uint) (dto.TeamDetailsResponse, error) {
	team, err := tuc.tr.GetTeamByID(teamID)

	if err != nil {
		return dto.TeamDetailsResponse{}, err
	}

	teamDetails := dto.TeamDetailsResponse{
		Name:        team.Name,
		Description: team.Description,
		Capacity:    team.Capacity,
	}

	for _, member := range team.TeamMembers {
		teamDetails.TeamMembers = append(teamDetails.TeamMembers, dto.TeamMemberResponse{
			UserID:   member.ID,
			Name:     member.User.Name,
			IsLeader: member.IsLeader,
		})
	}

	return teamDetails, nil
}
