package usecase

import (
	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/internal/team/entity"
	"github.com/alimikegami/compnouron/internal/team/repository"
)

type TeamUseCase struct {
	tr *repository.TeamRepository
}

func CreateNewTeamUseCase(tr *repository.TeamRepository) *TeamUseCase {
	return &TeamUseCase{tr: tr}
}

func (tuc *TeamUseCase) CreateTeam(userID uint, team dto.TeamRequest) error {
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

func (tuc *TeamUseCase) GetTeamsByUserID(userID uint) ([]dto.BriefTeamResponse, error) {
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

func (tuc *TeamUseCase) GetTeamDetailsByID(teamID uint) (dto.TeamDetailsResponse, error) {
	team, err := tuc.tr.GetTeamByID(teamID)

	if err != nil {
		return dto.TeamDetailsResponse{}, err
	}

	teamDetails := dto.TeamDetailsResponse{
		Name:        team.Name,
		Description: team.Description,
		Capacity:    team.Capacity,
	}

	members, err := tuc.tr.GetTeamMembersByID(teamID)
	if err != nil {
		return dto.TeamDetailsResponse{}, err
	}
	for _, member := range members {
		teamDetails.TeamMembers = append(teamDetails.TeamMembers, dto.TeamMemberResponse{
			UserID:   member.ID,
			Name:     member.User.Name,
			IsLeader: member.IsLeader,
		})
	}

	return teamDetails, nil
}
