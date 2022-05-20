package usecase

import (
	"testing"
	"time"

	competitionRepo "github.com/alimikegami/compnouron/internal/mocks/competition/repository"
	recruitmentRepo "github.com/alimikegami/compnouron/internal/mocks/recruitment/repository"
	userRepo "github.com/alimikegami/compnouron/internal/mocks/user/repository"
	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockRepo := userRepo.NewUserRepository(t)
	mockCompetition := competitionRepo.NewCompetitionRepository(t)
	mockRecruitment := recruitmentRepo.NewRecruitmentRepository(t)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetUserByEmail", "asdfa@gmail.com").Return(&entity.User{
			ID:                1,
			Name:              "Alim Ikegami",
			Email:             "asdfa@gmail.com",
			PhoneNumber:       "081111111111",
			SchoolInstitution: "Udayana University",
			Password:          "$2a$10$YefQPq3c5H7OalTHNFgo8Ob7Sxjc8F.fI3.ePvHOhOYCkqOGrFhm6",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		token, err := testUseCase.Login(&dto.Credential{
			Email:    "asdfa@gmail.com",
			Password: "asdfasfas",
		})
		assert.NoError(t, err)
		assert.NotNil(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockRepo.On("GetUserByEmail", "asdfa@gmail.com").Return(nil).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		token, err := testUseCase.Login(&dto.Credential{
			Email:    "asdfa@gmail.com",
			Password: "asdfasfas",
		})
		assert.Error(t, err)
		assert.Equal(t, token, "")
		mockRepo.AssertExpectations(t)
	})
}
<<<<<<< Updated upstream
=======

func TestGetCompetitionsData(t *testing.T) {
	mockRepo := userRepo.NewUserRepository(t)
	mockCompetition := competitionRepo.NewCompetitionRepository(t)
	mockRecruitment := recruitmentRepo.NewRecruitmentRepository(t)
	t.Run("success", func(t *testing.T) {
		mockCompetition.On("GetCompetitionByUserID", uint(1)).Return([]entityComp.Competition{
			{
				ID:                       1,
				Name:                     "Techoscape",
				Description:              "hackathon competition in Indonesia",
				ContactPerson:            "081239990127",
				IsTeam:                   1,
				IsTheSameInstitution:     1,
				RegistrationPeriodStatus: 1,
				TeamCapacity:             4,
				Level:                    "university student",
				CreatedAt:                time.Now(),
				UpdatedAt:                time.Now(),
				UserID:                   1,
			},
		}, nil).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		res, err := testUseCase.GetCompetitionsData(uint(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-error", func(t *testing.T) {
		mockCompetition.On("GetCompetitionByUserID", uint(1)).Return([]entityComp.Competition{}, errors.New("unexpected error")).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		res, err := testUseCase.GetCompetitionsData(uint(1))
		assert.Error(t, err)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetCompetitionRegistrationHistory(t *testing.T) {
	mockRepo := userRepo.NewUserRepository(t)
	mockCompetition := competitionRepo.NewCompetitionRepository(t)
	mockRecruitment := recruitmentRepo.NewRecruitmentRepository(t)
	t.Run("success", func(t *testing.T) {
		mockCompetition.On("GetCompetitionRegistrationByUserID", uint(1)).Return([]entityComp.CompetitionRegistration{
			{
				ID:               1,
				TeamID:           1,
				CompetitionID:    1,
				AcceptanceStatus: 1,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
				UserID:           1,
			},
		}, nil).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		res, err := testUseCase.GetCompetitionRegistrationHistory(uint(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unexpected-error", func(t *testing.T) {
		mockCompetition.On("GetCompetitionRegistrationByUserID", uint(1)).Return([]entityComp.CompetitionRegistration{}, errors.New("unexpected error")).Once()
		testUseCase := CreateNewUserUseCase(mockRepo, mockCompetition, mockRecruitment)
		res, err := testUseCase.GetCompetitionRegistrationHistory(uint(1))
		assert.Error(t, err)
		assert.Empty(t, res)
		mockRepo.AssertExpectations(t)
	})
}
>>>>>>> Stashed changes
