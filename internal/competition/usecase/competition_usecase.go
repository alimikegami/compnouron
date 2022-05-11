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

func (cuc *CompetitionUseCase) DeleteCompetition(competitionID uint, userID uint) error {
	competition, err := cuc.ur.GetCompetitionByID(competitionID)
	if err != nil {
		return err
	}
	if competition == (entity.Competition{}) {
		return errors.New("competition does not exist")
	}

	if competition.UserID != userID {
		return errors.New("action unauthorized")
	}
	err = cuc.ur.DeleteCompetition(competitionID)
	if err != nil {
		return err
	}

	return nil
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

func (cuc *CompetitionUseCase) GetCompetitionByID(id uint) (dto.CompetitionRequest, error) {
	competitionEntity, err := cuc.ur.GetCompetitionByID(id)
	if err != nil {
		return dto.CompetitionRequest{}, err
	}

	competitionRequest := dto.CompetitionRequest{
		Name:                 competitionEntity.Name,
		ContactPerson:        competitionEntity.ContactPerson,
		Description:          competitionEntity.Description,
		IsTheSameInstitution: competitionEntity.IsTheSameInstitution,
		IsTeam:               competitionEntity.IsTeam,
		TeamCapacity:         competitionEntity.TeamCapacity,
		Level:                competitionEntity.Level,
	}

	return competitionRequest, nil
}
