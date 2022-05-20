package usecase

import (
	"errors"
	"testing"

	recruitmentRepo "github.com/alimikegami/compnouron/internal/mocks/recruitment/repository"
	teamRepo "github.com/alimikegami/compnouron/internal/mocks/team/repository"
	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRecruitmentRepository struct {
	mock.Mock
}

func TestCreateRecruitment(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	req := dto.RecruitmentRequest{
		Role:        "Backend Engineer",
		Description: "Need Node.JS Developer",
		TeamID:      1,
	}
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("CreateRecruitment", entity.Recruitment{
			Role:                        "Backend Engineer",
			Description:                 "Need Node.JS Developer",
			TeamID:                      1,
			ApplicationAcceptanceStatus: 0,
		}).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitment(req, uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-create-recruitment-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("CreateRecruitment", entity.Recruitment{
			Role:                        "Backend Engineer",
			Description:                 "Need Node.JS Developer",
			TeamID:                      1,
			ApplicationAcceptanceStatus: 0,
		}).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitment(req, uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitment(req, uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitment(req, uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}

func TestUpdateRecruitment(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	req := dto.RecruitmentRequest{
		Role:        "Backend Engineer",
		Description: "Need Node.JS Developer",
		TeamID:      1,
	}
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("UpdateRecruitment", entity.Recruitment{
			ID:                          1,
			Role:                        "Backend Engineer",
			Description:                 "Need Node.JS Developer",
			TeamID:                      1,
			ApplicationAcceptanceStatus: 0,
		}).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.UpdateRecruitment(req, uint(1), uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-update-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("UpdateRecruitment", entity.Recruitment{
			ID:                          1,
			Role:                        "Backend Engineer",
			Description:                 "Need Node.JS Developer",
			TeamID:                      1,
			ApplicationAcceptanceStatus: 0,
		}).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.UpdateRecruitment(req, uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.UpdateRecruitment(req, uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.UpdateRecruitment(req, uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}
