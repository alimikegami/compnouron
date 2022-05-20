package usecase

import (
	"errors"
	"testing"
	"time"

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

func TestGetRecruitmentByID(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentByID", uint(1)).Return(entity.Recruitment{
			ID:                          1,
			Role:                        "Backend Engineer",
			Description:                 "Need Node.JS Developer",
			TeamID:                      1,
			ApplicationAcceptanceStatus: 0,
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Time{},
		}, nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		resp, err := testUseCase.GetRecruitmentByID(uint(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-recruitment-by-id-error", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentByID", uint(1)).Return(entity.Recruitment{}, errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		resp, err := testUseCase.GetRecruitmentByID(uint(1))
		assert.Error(t, err)
		assert.Empty(t, resp)
		mockRecuitmentRepo.AssertExpectations(t)
	})
}

func TestCreateRecruitmentApplication(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentApplicationByUserID", uint(1)).Return([]entity.RecruitmentApplication{}, nil).Once()
		mockRecuitmentRepo.On("CreateRecruitmentApplication", entity.RecruitmentApplication{
			UserID:           1,
			RecruitmentID:    1,
			AcceptanceStatus: 0,
		}).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}, uint(1))
		assert.NoError(t, err)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-create-recruitment-application-error", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentApplicationByUserID", uint(1)).Return([]entity.RecruitmentApplication{}, nil).Once()
		mockRecuitmentRepo.On("CreateRecruitmentApplication", entity.RecruitmentApplication{
			UserID:           1,
			RecruitmentID:    1,
			AcceptanceStatus: 0,
		}).Return(errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}, uint(1))
		assert.Error(t, err)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("recruitment-application-has-been-made", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentApplicationByUserID", uint(1)).Return([]entity.RecruitmentApplication{
			{
				ID:               1,
				UserID:           1,
				RecruitmentID:    1,
				AcceptanceStatus: 1,
			},
		}, nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}, uint(1))
		assert.Error(t, err)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("success-2", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentApplicationByUserID", uint(1)).Return([]entity.RecruitmentApplication{
			{
				ID:               1,
				UserID:           1,
				RecruitmentID:    1,
				AcceptanceStatus: 2,
			},
		}, nil).Once()
		mockRecuitmentRepo.On("CreateRecruitmentApplication", entity.RecruitmentApplication{
			UserID:           1,
			RecruitmentID:    1,
			AcceptanceStatus: 0,
		}).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}, uint(1))
		assert.NoError(t, err)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})
}

func TestRejectRecruitmentApplication(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("RejectRecruitmentApplication", uint(1)).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.RejectRecruitmentApplication(uint(1), uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-update-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("RejectRecruitmentApplication", uint(1)).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.RejectRecruitmentApplication(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.RejectRecruitmentApplication(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.RejectRecruitmentApplication(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}

func TestOpenRecruitmentApplicationPeriod(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("OpenRecruitmentApplicationPeriod", uint(1)).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.OpenRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-update-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("OpenRecruitmentApplicationPeriod", uint(1)).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.OpenRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.OpenRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.OpenRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}

func TestCloseRecruitmentApplicationPeriod(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("CloseRecruitmentApplicationPeriod", uint(1)).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CloseRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-update-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("CloseRecruitmentApplicationPeriod", uint(1)).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CloseRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CloseRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.CloseRecruitmentApplicationPeriod(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}

func TestDeleteRecruitmentByID(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("DeleteRecruitmentByID", uint(1)).Return(nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.DeleteRecruitmentByID(uint(1), uint(1))
		assert.NoError(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-update-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(1), nil).Once()
		mockRecuitmentRepo.On("DeleteRecruitmentByID", uint(1)).Return(errors.New("unxpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.DeleteRecruitmentByID(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(2), nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.DeleteRecruitmentByID(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-team-leader-error", func(t *testing.T) {
		mockTeamRepo.On("GetTeamLeader", uint(1)).Return(uint(0), errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		err := testUseCase.DeleteRecruitmentByID(uint(1), uint(1))
		assert.Error(t, err)
		mockTeamRepo.AssertExpectations(t)
	})
}

func TestGetRecruitmentByTeamID(t *testing.T) {
	mockRecuitmentRepo := recruitmentRepo.NewRecruitmentRepository(t)
	mockTeamRepo := teamRepo.NewTeamRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentByTeamID", uint(1)).Return([]entity.Recruitment{
			{
				ID:                          1,
				Role:                        "Backend Engineer",
				Description:                 "Need Node.JS Developer",
				TeamID:                      1,
				ApplicationAcceptanceStatus: 0,
				CreatedAt:                   time.Now(),
				UpdatedAt:                   time.Time{},
			},
			{
				ID:                          1,
				Role:                        "Frontend Engineer",
				Description:                 "Need React.JS Developer",
				TeamID:                      1,
				ApplicationAcceptanceStatus: 0,
				CreatedAt:                   time.Now(),
				UpdatedAt:                   time.Time{},
			},
		}, nil).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		resp, err := testUseCase.GetRecruitmentByTeamID(uint(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
		mockRecuitmentRepo.AssertExpectations(t)
		mockRecuitmentRepo.AssertExpectations(t)
	})

	t.Run("unexpected-get-recruitment-by-id-error", func(t *testing.T) {
		mockRecuitmentRepo.On("GetRecruitmentByTeamID", uint(1)).Return([]entity.Recruitment{}, errors.New("unexpected db error")).Once()
		testUseCase := CreateNewRecruitmentUseCase(mockRecuitmentRepo, mockTeamRepo)
		resp, err := testUseCase.GetRecruitmentByTeamID(uint(1))
		assert.Error(t, err)
		assert.Empty(t, resp)
		mockRecuitmentRepo.AssertExpectations(t)
	})
}
