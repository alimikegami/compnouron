package usecase

import (
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) GetUserByEmail(email string) *entity.User {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User)
}

func (mock *MockUserRepository) CreateUser(user *entity.User) error {
	args := mock.Called()

	return args.Error(0)
}

// func TestCreateUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)

// 	mockRepo.On("CreateUser").Return(nil)

// 	testUseCase := CreateNewUserUseCase(mockRepo)
// 	err := testUseCase.CreateUser(&entity.User{
// 		Name:              "Alim Ikegami",
// 		Email:             "sdafsfa@gmail.com",
// 		PhoneNumber:       "081111111111",
// 		Password:          "asdfasfas",
// 		SchoolInstitution: "Udayana University",
// 	})

// 	assert.Nil(t, err)
// 	mockRepo.AssertExpectations(t)
// }
