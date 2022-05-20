package usecase

import (
	"errors"
	"testing"

	"github.com/alimikegami/compnouron/internal/competition/entity"
	mockRepo "github.com/alimikegami/compnouron/internal/mocks/competition/repository"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCompetition(t *testing.T) {
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
		mockRepo.On("DeleteCompetition", uint(1)).Return(nil).Once()
		testUseCase := CreateNewCompetitionUseCase(mockRepo)
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
		testUseCase := CreateNewCompetitionUseCase(mockRepo)
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
		testUseCase := CreateNewCompetitionUseCase(mockRepo)
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
		testUseCase := CreateNewCompetitionUseCase(mockRepo)
		err := testUseCase.DeleteCompetition(uint(1), uint(3))
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
