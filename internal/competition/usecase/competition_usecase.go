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
	competition, err := cuc.ur.GetCompetitionByID(competitionID)
	if err != nil {
		return err
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

func (cuc *CompetitionUseCase) Register(competitionRegistration dto.CompetitionRegistrationRequest) error {
	comp, err := cuc.ur.GetCompetitionByID(competitionRegistration.CompetitionID)
	if err != nil {
		return err
	}
	if comp.IsOpen == 0 {
		return errors.New("registration period is over")
	}
	competitionRegistrationEntity := entity.CompetitionRegistration{
		UserID:        competitionRegistration.UserID,
		CompetitionID: competitionRegistration.CompetitionID,
		TeamID:        competitionRegistration.TeamID,
	}

	err = cuc.ur.Register(competitionRegistrationEntity)
	return err
}

func (cuc *CompetitionUseCase) RejectCompetitionRegistration(id uint) error {
	err := cuc.ur.RejectCompetitionRegistration(id)

	return err
}

func (cuc *CompetitionUseCase) AcceptCompetitionRegistration(id uint) error {
	err := cuc.ur.AcceptCompetitionRegistration(id)

	return err
}

func (cuc *CompetitionUseCase) OpenCompetitionRegistrationPeriod(id uint) error {
	err := cuc.ur.OpenCompetitionRegistrationPeriod(id)

	return err
}

func (cuc *CompetitionUseCase) CloseCompetitionRegistrationPeriod(id uint) error {
	err := cuc.ur.CloseCompetitionRegistrationPeriod(id)

	return err
}

func (cuc *CompetitionUseCase) GetCompetitionRegistration(id uint) (interface{}, error) {

	competition, err := cuc.ur.GetCompetitionByID(id)

	if err != nil {
		return nil, err
	}

	competitionRegistrations, err := cuc.ur.GetCompetitionRegistration(id)

	if err != nil {
		return nil, err
	}

	if competition.IsTeam == 1 {
		var competitionRegistrationsResponse []dto.TeamCompetitionRegistrationResponse
		for _, competitionRegistration := range competitionRegistrations {
			competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.TeamCompetitionRegistrationResponse{
				ID:            competitionRegistration.ID,
				TeamID:        competitionRegistration.TeamID,
				TeamName:      competitionRegistration.Team.Name,
				CompetitionID: competitionRegistration.CompetitionID,
				IsAccepted:    competitionRegistration.IsAccepted,
			})
		}

		return competitionRegistrationsResponse, nil
	}

	var competitionRegistrationsResponse []dto.IndividualCompetitionRegistrationResponse
	for _, competitionRegistration := range competitionRegistrations {
		competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.IndividualCompetitionRegistrationResponse{
			ID:                competitionRegistration.ID,
			UserID:            competitionRegistration.User.ID,
			UserName:          competitionRegistration.User.Name,
			PhoneNumber:       competitionRegistration.User.PhoneNumber,
			Email:             competitionRegistration.User.Email,
			SchoolInstitution: competitionRegistration.User.SchoolInstitution,
			CompetitionID:     competitionRegistration.CompetitionID,
			IsAccepted:        competitionRegistration.IsAccepted,
		})
	}

	return competitionRegistrationsResponse, nil
}
