package usecase

import (
	"errors"
	"testing"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	"github.com/alimikegami/compnouron/internal/competition/entity"
	mockRepo "github.com/alimikegami/compnouron/internal/mocks/competition/repository"
	teamRepo "github.com/alimikegami/compnouron/internal/mocks/team/repository"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCompetition(t *testing.T) {
	mockRepo := mockRepo.NewCompetitionRepository(t)
	teamRepository := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("DeleteCompetition", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.DeleteCompetition(uint(1), uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-delete-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("DeleteCompetition", uint(1)).Return(errors.New("errors db")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.DeleteCompetition(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   2,
		}, nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.DeleteCompetition(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-competition-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, errors.New("errors")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.DeleteCompetition(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestOpenCompetitionRegistrationPeriod(t *testing.T) {
	mockRepo := mockRepo.NewCompetitionRepository(t)
	teamRepository := teamRepo.NewTeamRepository(t)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("OpenCompetitionRegistrationPeriod", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.OpenCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-open-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("OpenCompetitionRegistrationPeriod", uint(1)).Return(errors.New("errors db")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.OpenCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   2,
		}, nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.OpenCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-competition-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, errors.New("errors")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.OpenCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCloseCompetitionRegistrationPeriod(t *testing.T) {
	teamRepository := teamRepo.NewTeamRepository(t)

	mockRepo := mockRepo.NewCompetitionRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("CloseCompetitionRegistrationPeriod", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CloseCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-close-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("CloseCompetitionRegistrationPeriod", uint(1)).Return(errors.New("errors db")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CloseCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   2,
		}, nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CloseCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-competition-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, errors.New("errors")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CloseCompetitionRegistrationPeriod(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestAcceptCompetitionRegistration(t *testing.T) {
	teamRepository := teamRepo.NewTeamRepository(t)

	mockRepo := mockRepo.NewCompetitionRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("AcceptCompetitionRegistration", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.AcceptCompetitionRegistration(uint(1), uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-accept-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("AcceptCompetitionRegistration", uint(1)).Return(errors.New("errors db")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.AcceptCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   2,
		}, nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.AcceptCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-competition-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, errors.New("errors")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.AcceptCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestRejectCompetitionRegistration(t *testing.T) {
	teamRepository := teamRepo.NewTeamRepository(t)
	mockRepo := mockRepo.NewCompetitionRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("RejectCompetitionRegistration", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.RejectCompetitionRegistration(uint(1), uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-reject-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, nil).Once()
		mockRepo.On("RejectCompetitionRegistration", uint(1)).Return(errors.New("errors db")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.RejectCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   2,
		}, nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.RejectCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-competition-error", func(t *testing.T) {
		mockRepo.On("GetCompetitionByID", uint(1)).Return(entity.Competition{
			ID:                       1,
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}, errors.New("errors")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.RejectCompetitionRegistration(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCrea(t *testing.T) {
	mockRepo := mockRepo.NewCompetitionRepository(t)
	teamRepository := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateCompetition", &entity.Competition{
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CreateCompetition(dto.CompetitionRequest{
			Name:                 "technoscape",
			Description:          "asdf",
			ContactPerson:        "081239990128",
			IsTeam:               1,
			IsTheSameInstitution: 1,
			TeamCapacity:         3,
			Level:                "Uni student",
		}, uint(3))
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateCompetition", &entity.Competition{
			Name:                     "technoscape",
			Description:              "asdf",
			ContactPerson:            "081239990128",
			IsTeam:                   1,
			IsTheSameInstitution:     1,
			RegistrationPeriodStatus: 0,
			TeamCapacity:             3,
			Level:                    "Uni student",
			UserID:                   3,
		}).Return(errors.New("db error")).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo, teamRepository)
		err := testUseCase.CreateCompetition(dto.CompetitionRequest{
			Name:                 "technoscape",
			Description:          "asdf",
			ContactPerson:        "081239990128",
			IsTeam:               1,
			IsTheSameInstitution: 1,
			TeamCapacity:         3,
			Level:                "Uni student",
		}, uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
