package usecase

import (
	"errors"

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
		Name:          competition.Name,
		Description:   competition.Description,
		ContactPerson: competition.ContactPerson,
		IsTeam:        competition.IsTeam,
		TeamCapacity:  competition.TeamCapacity,
		Level:         competition.Level,
		UserID:        userID,
	}
	err := cuc.ur.CreateCompetition(competitionEntity)
	return err
}

func (cuc *CompetitionUseCase) DeleteCompetition(competitionID uint, userID uint) error {
	competition := cuc.ur.GetCompetitionByID(competitionID)
	if competition == nil {
		return errors.New("competition does not exist")
	}

	if competition.UserID != userID {
		return errors.New("action unauthorized")
	}
	err := cuc.ur.DeleteCompetition(competitionID)
	if err != nil {
		return err
	}

	return nil
}

func (cuc *CompetitionUseCase) UpdateCompetition(competition dto.CompetitionRequest, id uint) error {
	competitionEntity := &entity.Competition{
		ID:            id,
		Name:          competition.Name,
		Description:   competition.Description,
		ContactPerson: competition.ContactPerson,
		IsTeam:        competition.IsTeam,
		TeamCapacity:  competition.TeamCapacity,
		Level:         competition.Level,
	}
	err := cuc.ur.UpdateCompetition(*competitionEntity)
	return err
}

func (cuc *CompetitionUseCase) Register(competitionRegistration dto.CompetitionRegistrationRequest) error {
	competitionRegistrationEntity := entity.CompetitionRegistration{
		UserID:        competitionRegistration.UserID,
		CompetitionID: competitionRegistration.CompetitionID,
		TeamID:        competitionRegistration.TeamID,
	}

	err := cuc.ur.Register(competitionRegistrationEntity)
	return err
}
