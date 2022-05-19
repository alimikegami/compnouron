package usecase

import (
	"fmt"

	compRepo "github.com/alimikegami/compnouron/internal/competition/repository"
	recRepo "github.com/alimikegami/compnouron/internal/recruitment/repository"

	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/alimikegami/compnouron/internal/user/repository"

	"github.com/alimikegami/compnouron/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(user *dto.UserRegistrationRequest) error
	Login(credential *dto.Credential) (string, error)
	GetCompetitionRegistrationHistory(userID uint) ([]dto.UserCompetitionHistory, error)
	GetRecruitmentApplicationHistory(userID uint) ([]dto.UserRecruitmentApplicationHistory, error)
}

type UserUseCaseImpl struct {
	ur repository.UserRepository
	cr compRepo.CompetitionRepository
	rr recRepo.RecruitmentRepository
}

func CreateNewUserUseCase(ur repository.UserRepository, cr compRepo.CompetitionRepository, rr recRepo.RecruitmentRepository) UserUseCase {
	return &UserUseCaseImpl{ur: ur, cr: cr, rr: rr}
}

func (us *UserUseCaseImpl) CreateUser(user *dto.UserRegistrationRequest) error {
	var skills []entity.Skill

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	userEntity := entity.User{
		Name:              user.Name,
		Email:             user.Email,
		Password:          string(hash),
		PhoneNumber:       user.PhoneNumber,
		SchoolInstitution: user.SchoolInstitution,
	}

	userID, err := us.ur.CreateUser(userEntity)
	if err != nil {
		return err
	}

	for _, skill := range user.Skills {
		skills = append(skills, entity.Skill{
			Name:   skill.Name,
			UserID: userID,
		})
	}

	err = us.ur.AddUserSkills(skills)
	return err
}

func (us *UserUseCaseImpl) Login(credential *dto.Credential) (string, error) {
	user := us.ur.GetUserByEmail(credential.Email)
	if user == nil {
		return "", fmt.Errorf("credentials dont match")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		return "", fmt.Errorf("credentials dont match")
	}
	token, err := utils.CreateSignedJWTToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return token, nil
}

func (us *UserUseCaseImpl) GetCompetitionRegistrationHistory(userID uint) ([]dto.UserCompetitionHistory, error) {
	var history []dto.UserCompetitionHistory
	comps, err := us.cr.GetCompetitionRegistrationByUserID(userID)
	for _, comp := range comps {
		history = append(history, dto.UserCompetitionHistory{
			CompetitionRegistrationID: comp.ID,
			AcceptanceStatus:          comp.AcceptanceStatus,
			CompetitionName:           comp.Competition.Name,
			CompetitionID:             comp.CompetitionID,
		})
	}
	return history, err
}

func (us *UserUseCaseImpl) GetRecruitmentApplicationHistory(userID uint) ([]dto.UserRecruitmentApplicationHistory, error) {
	var history []dto.UserRecruitmentApplicationHistory
	recs, err := us.rr.GetRecruitmentApplicationByUserID(userID)
	for _, rec := range recs {
		history = append(history, dto.UserRecruitmentApplicationHistory{
			RecruitmentApplicationID: rec.ID,
			RecruitmentID:            rec.RecruitmentID,
			RecruitmentRole:          rec.Recruitment.Role,
			AcceptanceStatus:         uint(rec.AcceptanceStatus),
		})
	}
	return history, err
}
