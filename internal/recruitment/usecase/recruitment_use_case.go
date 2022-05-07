package usecase

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/internal/recruitment/repository"
)

type RecruitmentUseCase struct {
	rr *repository.RecruitmentRepository
}

func CreateNewRecruitmentUseCase(rr *repository.RecruitmentRepository) *RecruitmentUseCase {
	return &RecruitmentUseCase{rr: rr}
}

func (ruc *RecruitmentUseCase) CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error {
	recruitmentEntity := entity.Recruitment{
		Role:        recruitmentRequest.Role,
		Description: recruitmentRequest.Description,
		TeamID:      recruitmentRequest.TeamID,
	}
	fmt.Println(recruitmentEntity)
	err := ruc.rr.CreateRecruitment(recruitmentEntity)
	return err
}
