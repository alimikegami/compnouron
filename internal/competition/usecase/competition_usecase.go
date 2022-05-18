package usecase

import (
	"errors"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	"github.com/alimikegami/compnouron/internal/competition/entity"
	"github.com/alimikegami/compnouron/internal/competition/repository"
)

type CompetitionUseCaseImpl struct {
	ur repository.CompetitionRepository
}

type CompetitionUseCase interface {
	CreateCompetition(competition dto.CompetitionRequest, userID uint) error
	DeleteCompetition(competitionID uint, userID uint) error
	UpdateCompetition(competition dto.CompetitionRequest, id uint, userID uint) error
	GetCompetitions(limit int, offset int) ([]dto.CompetitionResponse, error)
	Register(competitionRegistration dto.CompetitionRegistrationRequest) error
	RejectCompetitionRegistration(id uint) error
	AcceptCompetitionRegistration(id uint) error
	OpenCompetitionRegistrationPeriod(id uint) error
	CloseCompetitionRegistrationPeriod(id uint) error
	GetCompetitionRegistration(id uint) (interface{}, error)
	GetAcceptedCompetitionParticipants(id uint) (interface{}, error)
	SearchCompetition(limit int, offset int, keyword string) ([]dto.CompetitionResponse, error)
}

func CreateNewCompetitionUseCase(ur repository.CompetitionRepository) CompetitionUseCase {
	return &CompetitionUseCaseImpl{ur: ur}
}

func (cuc *CompetitionUseCaseImpl) CreateCompetition(competition dto.CompetitionRequest, userID uint) error {
	competitionEntity := &entity.Competition{
		Name:                 competition.Name,
		Description:          competition.Description,
		ContactPerson:        competition.ContactPerson,
		IsTeam:               competition.IsTeam,
		IsTheSameInstitution: competition.IsTheSameInstitution,
		TeamCapacity:         competition.TeamCapacity,
		Level:                competition.Level,
		UserID:               userID,
	}
	err := cuc.ur.CreateCompetition(competitionEntity)
	return err
}

func (cuc *CompetitionUseCaseImpl) DeleteCompetition(competitionID uint, userID uint) error {
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

func (cuc *CompetitionUseCaseImpl) UpdateCompetition(competition dto.CompetitionRequest, id uint, userID uint) error {
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

	competitionData, err := cuc.ur.GetCompetitionByID(id)
	if err != nil {
		return err
	}
	if competitionData.ID != userID {
		return errors.New("action unauthorized")
	}

	err = cuc.ur.UpdateCompetition(*competitionEntity)
	return err
}

func (cuc *CompetitionUseCaseImpl) GetCompetitions(limit int, offset int) ([]dto.CompetitionResponse, error) {
	var competitionsResponse []dto.CompetitionResponse
	competitionsEntity, err := cuc.ur.GetCompetitions(limit, offset)
	if err != nil {
		return []dto.CompetitionResponse{}, err
	}

	for _, competition := range competitionsEntity {
		competitionsResponse = append(competitionsResponse, dto.CompetitionResponse{
			ID:            competition.ID,
			Name:          competition.Name,
			ContactPerson: competition.ContactPerson,
			IsTeam:        competition.IsTeam,
			Level:         competition.Level,
		})
	}

	return competitionsResponse, nil
}

func (cuc *CompetitionUseCaseImpl) Register(competitionRegistration dto.CompetitionRegistrationRequest) error {
	comp, err := cuc.ur.GetCompetitionByID(competitionRegistration.CompetitionID)
	if err != nil {
		return err
	}
	if comp.RegistrationPeriodStatus == 0 {
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

func (cuc *CompetitionUseCaseImpl) RejectCompetitionRegistration(id uint) error {
	err := cuc.ur.RejectCompetitionRegistration(id)

	return err
}

func (cuc *CompetitionUseCaseImpl) AcceptCompetitionRegistration(id uint) error {
	err := cuc.ur.AcceptCompetitionRegistration(id)

	return err
}

func (cuc *CompetitionUseCaseImpl) OpenCompetitionRegistrationPeriod(id uint) error {
	err := cuc.ur.OpenCompetitionRegistrationPeriod(id)

	return err
}

func (cuc *CompetitionUseCaseImpl) CloseCompetitionRegistrationPeriod(id uint) error {
	err := cuc.ur.CloseCompetitionRegistrationPeriod(id)

	return err
}

func (cuc *CompetitionUseCaseImpl) GetCompetitionRegistration(id uint) (interface{}, error) {

	competition, err := cuc.ur.GetCompetitionRegistration(id)

	if err != nil {
		return nil, err
	}

	if competition.IsTeam == 1 {
		var competitionRegistrationsResponse []dto.TeamCompetitionRegistrationResponse
		for _, competitionRegistration := range competition.CompetitionRegistrations {
			competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.TeamCompetitionRegistrationResponse{
				ID:               competitionRegistration.ID,
				TeamID:           competitionRegistration.TeamID,
				TeamName:         competitionRegistration.Team.Name,
				CompetitionID:    competitionRegistration.CompetitionID,
				AcceptanceStatus: competitionRegistration.AcceptanceStatus,
			})
		}

		return competitionRegistrationsResponse, nil
	}

	var competitionRegistrationsResponse []dto.IndividualCompetitionRegistrationResponse
	for _, competitionRegistration := range competition.CompetitionRegistrations {
		competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.IndividualCompetitionRegistrationResponse{
			ID:                competitionRegistration.ID,
			UserID:            competitionRegistration.User.ID,
			UserName:          competitionRegistration.User.Name,
			PhoneNumber:       competitionRegistration.User.PhoneNumber,
			Email:             competitionRegistration.User.Email,
			SchoolInstitution: competitionRegistration.User.SchoolInstitution,
			CompetitionID:     competitionRegistration.CompetitionID,
			AcceptanceStatus:  competitionRegistration.AcceptanceStatus,
		})
	}

	return competitionRegistrationsResponse, nil
}

func (cuc *CompetitionUseCaseImpl) GetAcceptedCompetitionParticipants(id uint) (interface{}, error) {

	competition, err := cuc.ur.GetAcceptedCompetitionParticipants(id)

	if err != nil {
		return nil, err
	}

	if competition.IsTeam == 1 {
		var competitionRegistrationsResponse []dto.TeamCompetitionRegistrationResponse
		for _, competitionRegistration := range competition.CompetitionRegistrations {
			competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.TeamCompetitionRegistrationResponse{
				ID:               competitionRegistration.ID,
				TeamID:           competitionRegistration.TeamID,
				TeamName:         competitionRegistration.Team.Name,
				CompetitionID:    competitionRegistration.CompetitionID,
				AcceptanceStatus: competitionRegistration.AcceptanceStatus,
			})
		}

		return competitionRegistrationsResponse, nil
	}

	var competitionRegistrationsResponse []dto.IndividualCompetitionRegistrationResponse
	for _, competitionRegistration := range competition.CompetitionRegistrations {
		competitionRegistrationsResponse = append(competitionRegistrationsResponse, dto.IndividualCompetitionRegistrationResponse{
			ID:                competitionRegistration.ID,
			UserID:            competitionRegistration.User.ID,
			UserName:          competitionRegistration.User.Name,
			PhoneNumber:       competitionRegistration.User.PhoneNumber,
			Email:             competitionRegistration.User.Email,
			SchoolInstitution: competitionRegistration.User.SchoolInstitution,
			CompetitionID:     competitionRegistration.CompetitionID,
			AcceptanceStatus:  competitionRegistration.AcceptanceStatus,
		})
	}

	return competitionRegistrationsResponse, nil
}

func (cuc *CompetitionUseCaseImpl) SearchCompetition(limit int, offset int, keyword string) ([]dto.CompetitionResponse, error) {
	var competitionsResponse []dto.CompetitionResponse
	competitions, err := cuc.ur.SearchCompetition(limit, offset, keyword)
	if err != nil {
		return []dto.CompetitionResponse{}, err
	}

	for _, competition := range competitions {
		competitionsResponse = append(competitionsResponse, dto.CompetitionResponse{
			ID:            competition.ID,
			Name:          competition.Name,
			ContactPerson: competition.ContactPerson,
			IsTeam:        competition.IsTeam,
			Level:         competition.Level,
		})
	}

	return competitionsResponse, nil
}
