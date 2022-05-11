package usecase

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/internal/recruitment/repository"
)

type RecruitmentUseCase interface {
	CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error
	UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint) error
	CreateRecruitmentApplication(recruitmentApplication dto.RecruitmentApplicationRequest, userID uint) error
	GetRecruitmentDetailsByID(id uint) (dto.RecruitmentDetailsResponse, error)
	GetRecruitmentByUserID(id uint) (dto.RecruitmentsResponse, error)
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

func (ruc *RecruitmentUseCaseImpl) CreateRecruitmentApplication(recruitmentApplicationRequest dto.RecruitmentApplicationRequest, userID uint) error {
	recruitmentApplicationEntity := entity.RecruitmentApplication{
		UserID:        userID,
		RecruitmentID: recruitmentApplicationRequest.RecruitmentID,
		IsAccepted:    0,
	}

	err := ruc.rr.CreateRecruitmentApplication(recruitmentApplicationEntity)

	return err
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitmentDetailsByID(id uint) (dto.RecruitmentDetailsResponse, error) {
	var recruitmentApplicationsResponse []dto.RecruitmentApplicationResponse
	recruitment, err := ruc.rr.GetRecruitmentByID(id)
	if err != nil {
		return dto.RecruitmentDetailsResponse{}, err
	}

	recruitmentApplications, err := ruc.rr.GetRecruitmentApplicationByRecruitmentID(id)
	if err != nil {
		return dto.RecruitmentDetailsResponse{}, err
	}

	recruitmentResponse := dto.RecruitmentResponse{
		ID:          recruitment.ID,
		Role:        recruitment.Role,
		Description: recruitment.Description,
		TeamID:      recruitment.TeamID,
		TeamName:    recruitment.Team.Name,
	}

	for _, recruitmentApplication := range recruitmentApplications {
		recruitmentApplicationsResponse = append(recruitmentApplicationsResponse, dto.RecruitmentApplicationResponse{
			ID:            recruitmentApplication.ID,
			UserID:        recruitmentApplication.UserID,
			RecruitmentID: recruitmentApplication.RecruitmentID,
			IsAccepted:    recruitmentApplication.IsAccepted,
			UserName:      recruitmentApplication.User.Name,
		})
	}

	return dto.RecruitmentDetailsResponse{
		Recruitment:             recruitmentResponse,
		RecruitmentApplications: recruitmentApplicationsResponse,
	}, nil
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitmentByUserID(id uint) (dto.RecruitmentsResponse, error) {
	var recruitments dto.RecruitmentsResponse
	result, err := ruc.rr.GetRecruitmentByUserID(id)
	fmt.Println(result)
	for _, recruitment := range result {
		recruitments = append(recruitments, dto.RecruitmentResponse{
			ID:          recruitment.ID,
			Role:        recruitment.Role,
			Description: recruitment.Description,
			TeamID:      recruitment.TeamID,
			TeamName:    recruitment.Team.Name,
		})
	}

	return recruitments, err
}
