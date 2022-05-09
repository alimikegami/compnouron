package usecase

import (
	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/internal/recruitment/repository"
)

type RecruitmentUseCase interface {
	CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error
	UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint) error
}

type RecruitmentUseCaseImpl struct {
	rr repository.RecruitmentRepository
}

func CreateNewRecruitmentUseCase(rr repository.RecruitmentRepository) RecruitmentUseCase {
	return &RecruitmentUseCaseImpl{rr: rr}
}

func (ruc *RecruitmentUseCaseImpl) CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error {
	recruitmentEntity := entity.Recruitment{
		Role:        recruitmentRequest.Role,
		Description: recruitmentRequest.Description,
		TeamID:      recruitmentRequest.TeamID,
	}
	err := ruc.rr.CreateRecruitment(recruitmentEntity)
	return err
}

func (ruc *RecruitmentUseCaseImpl) UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint) error {
	recruitmentEntity := entity.Recruitment{
		ID:          id,
		Role:        recruitmentRequest.Role,
		Description: recruitmentRequest.Description,
		TeamID:      recruitmentRequest.TeamID,
	}

	err := ruc.rr.UpdateRecruitment(recruitmentEntity)
	return err
}
