package usecase

import (
	"testing"
	"time"

	userRepo "github.com/alimikegami/compnouron/internal/mocks/user/repository"
	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockRepo := userRepo.NewUserRepository(t)
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
		testUseCase := CreateNewUserUseCase(mockRepo)
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
		testUseCase := CreateNewUserUseCase(mockRepo)
		token, err := testUseCase.Login(&dto.Credential{
			Email:    "asdfa@gmail.com",
			Password: "asdfasfas",
		})
		assert.Error(t, err)
		assert.Equal(t, token, "")
		mockRepo.AssertExpectations(t)
	})
}
