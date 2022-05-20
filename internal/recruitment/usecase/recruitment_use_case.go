package usecase

import (
	"errors"

	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/internal/recruitment/repository"
	teamRepository "github.com/alimikegami/compnouron/internal/team/repository"
)

type RecruitmentUseCase interface {
	CreateRecruitment(recruitmentRequest dto.RecruitmentRequest, userID uint) error
	UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint, userID uint) error
	GetRecruitmentByID(id uint) (dto.RecruitmentResponse, error)
	CreateRecruitmentApplication(recruitmentApplication dto.RecruitmentApplicationRequest, userID uint) error
	GetRecruitmentDetailsByID(id uint, userID uint) (dto.RecruitmentDetailsResponse, error)
	GetRecruitmentByTeamID(id uint) (dto.RecruitmentsResponse, error)
	RejectRecruitmentApplication(id uint, userID uint) error
	AcceptRecruitmentApplication(id uint, userID uint) error
	DeleteRecruitmentByID(id uint, userID uint) error
	OpenRecruitmentApplicationPeriod(id uint, userID uint) error
	CloseRecruitmentApplicationPeriod(id uint, userID uint) error
	GetRecruitments(limit int, offset int) ([]dto.BriefRecruitmentResponse, error)
	SearchRecruitment(limit int, offset int, keyword string) ([]dto.BriefRecruitmentResponse, error)
}

type RecruitmentUseCaseImpl struct {
	rr repository.RecruitmentRepository
	tr teamRepository.TeamRepository
}

func CreateNewRecruitmentUseCase(rr repository.RecruitmentRepository, tr teamRepository.TeamRepository) RecruitmentUseCase {
	return &RecruitmentUseCaseImpl{rr: rr, tr: tr}
}

func (ruc *RecruitmentUseCaseImpl) CreateRecruitment(recruitmentRequest dto.RecruitmentRequest, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(recruitmentRequest.TeamID)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	recruitmentEntity := entity.Recruitment{
		Role:                        recruitmentRequest.Role,
		Description:                 recruitmentRequest.Description,
		TeamID:                      recruitmentRequest.TeamID,
		ApplicationAcceptanceStatus: 0,
	}
	err = ruc.rr.CreateRecruitment(recruitmentEntity)
	return err
}

func (ruc *RecruitmentUseCaseImpl) UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	recruitmentEntity := entity.Recruitment{
		ID:          id,
		Role:        recruitmentRequest.Role,
		Description: recruitmentRequest.Description,
		TeamID:      recruitmentRequest.TeamID,
	}

	err = ruc.rr.UpdateRecruitment(recruitmentEntity)
	return err
}

func (ruc *RecruitmentUseCaseImpl) CreateRecruitmentApplication(recruitmentApplicationRequest dto.RecruitmentApplicationRequest, userID uint) error {
	recruitmentApplicationEntity := entity.RecruitmentApplication{
		UserID:           userID,
		RecruitmentID:    recruitmentApplicationRequest.RecruitmentID,
		AcceptanceStatus: 0,
	}

	rec, err := ruc.rr.GetRecruitmentByID(recruitmentApplicationRequest.RecruitmentID)
	if err != nil {
		return err
	}

	if rec.ApplicationAcceptanceStatus == 0 || rec.ApplicationAcceptanceStatus == 2 {
		return errors.New("application is not opened yet")

	}
	teamMembers, err := ruc.tr.GetTeamByID(rec.TeamID)
	if err != nil {
		return err
	}
	for _, member := range teamMembers.TeamMembers {
		if member.UserID == userID {
			return errors.New("you are a team member")
		}
	}

	recHis, err := ruc.rr.GetRecruitmentApplicationByUserID(userID)
	if err != nil {
		return errors.New("internal server error")
	}
	for _, rec := range recHis {
		if rec.RecruitmentID == recruitmentApplicationRequest.RecruitmentID && rec.UserID == userID && (rec.AcceptanceStatus == 1 || rec.AcceptanceStatus == 0) {
			return errors.New("you have registered")
		}
	}
	err = ruc.rr.CreateRecruitmentApplication(recruitmentApplicationEntity)

	return err
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitments(limit int, offset int) ([]dto.BriefRecruitmentResponse, error) {
	var briefRecruitmentResponse []dto.BriefRecruitmentResponse
	recruitmentsEntity, err := ruc.rr.GetRecruitments(limit, offset)
	if err != nil {
		return []dto.BriefRecruitmentResponse{}, err
	}

	for _, recruitment := range recruitmentsEntity {
		briefRecruitmentResponse = append(briefRecruitmentResponse, dto.BriefRecruitmentResponse{
			ID:                          recruitment.ID,
			Role:                        recruitment.Role,
			TeamName:                    recruitment.Team.Name,
			ApplicationAcceptanceStatus: recruitment.ApplicationAcceptanceStatus,
		})
	}

	return briefRecruitmentResponse, nil
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitmentDetailsByID(id uint, userID uint) (dto.RecruitmentDetailsResponse, error) {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return dto.RecruitmentDetailsResponse{}, errors.New("internal server error")
	}

	if teamOwner != userID {
		return dto.RecruitmentDetailsResponse{}, errors.New("action unauthorized")
	}

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
		ID:                          recruitment.ID,
		Role:                        recruitment.Role,
		Description:                 recruitment.Description,
		TeamID:                      recruitment.TeamID,
		TeamName:                    recruitment.Team.Name,
		ApplicationAcceptanceStatus: recruitment.ApplicationAcceptanceStatus,
	}

	for _, recruitmentApplication := range recruitmentApplications {
		recruitmentApplicationsResponse = append(recruitmentApplicationsResponse, dto.RecruitmentApplicationResponse{
			ID:               recruitmentApplication.ID,
			UserID:           recruitmentApplication.UserID,
			RecruitmentID:    recruitmentApplication.RecruitmentID,
			AcceptanceStatus: recruitmentApplication.AcceptanceStatus,
			UserName:         recruitmentApplication.User.Name,
		})
	}

	return dto.RecruitmentDetailsResponse{
		Recruitment:             recruitmentResponse,
		RecruitmentApplications: recruitmentApplicationsResponse,
	}, nil
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitmentByTeamID(id uint) (dto.RecruitmentsResponse, error) {
	var recruitments dto.RecruitmentsResponse
	result, err := ruc.rr.GetRecruitmentByTeamID(id)
	for _, recruitment := range result {
		recruitments = append(recruitments, dto.RecruitmentResponse{
			ID:                          recruitment.ID,
			Role:                        recruitment.Role,
			Description:                 recruitment.Description,
			TeamID:                      recruitment.TeamID,
			TeamName:                    recruitment.Team.Name,
			ApplicationAcceptanceStatus: recruitment.ApplicationAcceptanceStatus,
		})
	}

	return recruitments, err
}

func (ruc *RecruitmentUseCaseImpl) RejectRecruitmentApplication(id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}
	err = ruc.rr.RejectRecruitmentApplication(id)

	return err
}

func (ruc *RecruitmentUseCaseImpl) AcceptRecruitmentApplication(id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	recruitmentApplication, err := ruc.rr.GetRecruitmentApplicationByID(id)
	if err != nil {
		return err
	}

	team, err := ruc.tr.GetTeamByID(recruitmentApplication.Recruitment.TeamID)
	if int(team.Capacity) <= len(team.TeamMembers) {
		return errors.New("The team is full")
	}

	err = ruc.rr.AcceptRecruitmentApplication(id)
	if err != nil {
		return err
	}

	err = ruc.tr.AddTeamMember(recruitmentApplication.UserID, recruitmentApplication.Recruitment.TeamID, 0)
	if err != nil {
		return err
	}
	return nil
}

func (ruc *RecruitmentUseCaseImpl) SearchRecruitment(limit int, offset int, keyword string) ([]dto.BriefRecruitmentResponse, error) {
	var recruitmentsResponse []dto.BriefRecruitmentResponse
	recruitments, err := ruc.rr.SearchRecruitment(limit, offset, keyword)
	if err != nil {
		return []dto.BriefRecruitmentResponse{}, err
	}

	for _, recruitment := range recruitments {
		recruitmentsResponse = append(recruitmentsResponse, dto.BriefRecruitmentResponse{
			ID:                          recruitment.ID,
			TeamName:                    recruitment.Team.Name,
			Role:                        recruitment.Role,
			ApplicationAcceptanceStatus: recruitment.ApplicationAcceptanceStatus,
		})
	}

	return recruitmentsResponse, nil
}

func (ruc *RecruitmentUseCaseImpl) GetRecruitmentByID(id uint) (dto.RecruitmentResponse, error) {
	recruitment, err := ruc.rr.GetRecruitmentByID(id)
	if err != nil {
		return dto.RecruitmentResponse{}, err
	}

	recruitmentResponse := dto.RecruitmentResponse{
		ID:                          recruitment.ID,
		Role:                        recruitment.Role,
		Description:                 recruitment.Description,
		TeamID:                      recruitment.TeamID,
		TeamName:                    recruitment.Team.Name,
		ApplicationAcceptanceStatus: recruitment.ApplicationAcceptanceStatus,
	}

	return recruitmentResponse, nil
}

func (ruc *RecruitmentUseCaseImpl) DeleteRecruitmentByID(id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	err = ruc.rr.DeleteRecruitmentByID(id)

	return err
}

func (ruc *RecruitmentUseCaseImpl) OpenRecruitmentApplicationPeriod(id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	err = ruc.rr.OpenRecruitmentApplicationPeriod(id)

	return err
}

func (ruc *RecruitmentUseCaseImpl) CloseRecruitmentApplicationPeriod(id uint, userID uint) error {
	teamOwner, err := ruc.tr.GetTeamLeader(id)
	if err != nil {
		return errors.New("internal server error")
	}

	if teamOwner != userID {
		return errors.New("action unauthorized")
	}

	err = ruc.rr.CloseRecruitmentApplicationPeriod(id)

	return err
}
