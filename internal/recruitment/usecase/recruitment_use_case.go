package usecase

import (
	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/internal/recruitment/repository"
	teamRepository "github.com/alimikegami/compnouron/internal/team/repository"
)

type RecruitmentUseCase interface {
	CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error
	UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint) error
	CreateRecruitmentApplication(recruitmentApplication dto.RecruitmentApplicationRequest, userID uint) error
	GetRecruitmentDetailsByID(id uint) (dto.RecruitmentDetailsResponse, error)
	GetRecruitmentByUserID(id uint) (dto.RecruitmentsResponse, error)
	RejectRecruitmentApplication(id uint) error
	AcceptRecruitmentApplication(id uint) error
	DeleteRecruitmentByID(id uint) error
}

type RecruitmentUseCaseImpl struct {
	rr repository.RecruitmentRepository
	tr teamRepository.TeamRepository
}

func CreateNewRecruitmentUseCase(rr repository.RecruitmentRepository, tr teamRepository.TeamRepository) RecruitmentUseCase {
	return &RecruitmentUseCaseImpl{rr: rr, tr: tr}
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
		IsRejected:    0,
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
			IsRejected:    recruitmentApplication.IsRejected,
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

func (ruc *RecruitmentUseCaseImpl) RejectRecruitmentApplication(id uint) error {
	err := ruc.rr.RejectRecruitmentApplication(id)

	return err
}

func (ruc *RecruitmentUseCaseImpl) AcceptRecruitmentApplication(id uint) error {
	err := ruc.rr.AcceptRecruitmentApplication(id)
	if err != nil {
		return err
	}

	recruitmentApplication, err := ruc.rr.GetRecruitmentApplicationByID(id)
	if err != nil {
		return err
	}
	err = ruc.tr.AddTeamMember(recruitmentApplication.UserID, recruitmentApplication.Recruitment.TeamID, 0)
	if err != nil {
		return err
	}
	return nil
}

func (ruc *RecruitmentUseCaseImpl) DeleteRecruitmentByID(id uint) error {
	err := ruc.rr.DeleteRecruitmentByID(id)

	return err
}
