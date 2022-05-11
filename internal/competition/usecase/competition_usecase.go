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

func (cuc *CompetitionUseCase) GetCompetitions(limit int, offset int) ([]dto.CompetitionResponse, error) {
	var competitionsResponse []dto.CompetitionResponse
	competitionsEntity, err := cuc.ur.GetCompetitions(limit, offset)
	if err != nil {
		return []dto.CompetitionResponse{}, err
	}

	for _, competition := range competitionsEntity {
		competitionsResponse = append(competitionsResponse, dto.CompetitionResponse{
			Name:          competition.Name,
			ContactPerson: competition.ContactPerson,
			IsTeam:        competition.IsTeam,
			Level:         competition.Level,
		})
	}

	return competitionsResponse, nil
}
