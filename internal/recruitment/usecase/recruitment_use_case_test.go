package usecase

import (
	"errors"
	"testing"

	mocks "github.com/alimikegami/compnouron/internal/mocks/recruitment/repository"
	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRecruitmentRepository struct {
	mock.Mock
}

func TestSuccessfulCreateRecruitmentApplication(t *testing.T) {
	mockRepo := mocks.NewRecruitmentRepository(t)
	mockRepo.On("CreateRecruitmentApplication", entity.RecruitmentApplication{
		UserID:        1,
		RecruitmentID: 1,
		IsAccepted:    0,
		IsRejected:    0,
	}).Return(nil)
	testUseCase := CreateNewRecruitmentUseCase(mockRepo)
	err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
		RecruitmentID: 1,
	}, 1)

	assert.Nil(t, err, "No error")
	mockRepo.AssertExpectations(t)
}

func TestInvalidCreateRecruitmentApplicationForeignKey(t *testing.T) {
	mockRepo := mocks.NewRecruitmentRepository(t)
	mockRepo.On("CreateRecruitmentApplication", entity.RecruitmentApplication{
		UserID:        1,
		RecruitmentID: 2,
		IsAccepted:    0,
		IsRejected:    0,
	}).Return(errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`compnouron`.`recruitment_applications`, CONSTRAINT `fk_recruitment_applications_recruitment` FOREIGN KEY (`recruitment_id`) REFERENCES `recruitments` (`id`))"))
	testUseCase := CreateNewRecruitmentUseCase(mockRepo)
	err := testUseCase.CreateRecruitmentApplication(dto.RecruitmentApplicationRequest{
		RecruitmentID: 2,
	}, 1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
