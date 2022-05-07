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
