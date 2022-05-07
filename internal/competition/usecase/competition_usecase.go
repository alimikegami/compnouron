package usecase

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	"github.com/alimikegami/compnouron/internal/competition/entity"
	"github.com/alimikegami/compnouron/internal/competition/repository"
)

type CompetitionUseCase struct {
	ur *repository.CompetitionRepository
}

func CreateNewCompetitionUseCase(ur *repository.CompetitionRepository) *CompetitionUseCase {
	return &CompetitionUseCase{ur: ur}
}

func (cuc *CompetitionUseCase) CreateCompetition(competition dto.CompetitionRequest, userID uint) error {
	competitionEntity := &entity.Competition{
		Name:                 competition.Name,
		Description:          competition.Description,
		ContactPerson:        competition.ContactPerson,
		IsTheSameInstitution: competition.IsTheSameInstitution,
		IsTeam:               competition.IsTeam,
		TeamCapacity:         competition.TeamCapacity,
		Level:                competition.Level,
		UserID:               userID,
	}
	err := cuc.ur.CreateCompetition(competitionEntity)
	return err
}

func (cuc *CompetitionUseCase) DeleteCompetition(competitionID uint, userID uint) {
	competition := cuc.ur.GetCompetitionByID(competitionID)
	if competition == nil {
		fmt.Println("competition does not exist")
		return
	}

	if competition.UserID == userID {
		cuc.ur.DeleteCompetition(competitionID)
	}
}

func (cuc *CompetitionUseCase) UpdateCompetition(competition dto.CompetitionRequest, id uint) error {
	competitionEntity := &entity.Competition{
		ID:                   id,
		Name:                 competition.Name,
		Description:          competition.Description,
		ContactPerson:        competition.ContactPerson,
		IsTheSameInstitution: competition.IsTheSameInstitution,
		IsTeam:               competition.IsTeam,
		TeamCapacity:         competition.TeamCapacity,
		Level:                competition.Level,
	}
	err := cuc.ur.UpdateCompetition(*competitionEntity)
	return err
}
