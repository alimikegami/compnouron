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
		UserID:      userID,
	}

	err := tuc.tr.CreateTeam(teamEntity)
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
