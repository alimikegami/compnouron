package repository

import (
	"errors"
	"fmt"

	"github.com/alimikegami/compnouron/db/pagination"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecruitmentRepository interface {
	CreateRecruitment(recruitment entity.Recruitment) error
	UpdateRecruitment(recruitment entity.Recruitment) error
	CreateRecruitmentApplication(recruitmentApplication entity.RecruitmentApplication) error
	GetRecruitmentByID(id uint) (entity.Recruitment, error)
	GetRecruitmentApplicationByRecruitmentID(id uint) ([]entity.RecruitmentApplication, error)
	GetRecruitmentByUserID(id uint) ([]entity.Recruitment, error)
	RejectRecruitmentApplication(id uint) error
	GetRecruitmentApplicationByID(id uint) (entity.RecruitmentApplication, error)
	AcceptRecruitmentApplication(id uint) error
	DeleteRecruitmentByID(id uint) error
	OpenRecruitmentApplicationPeriod(id uint) error
	CloseRecruitmentApplicationPeriod(id uint) error
	GetRecruitments(limit int, offset int) ([]entity.Recruitment, error)
}

type RecruitmentRepositoryImpl struct {
	db *gorm.DB
}

func CreateNewRecruitmentRepository(db *gorm.DB) RecruitmentRepository {
	return &RecruitmentRepositoryImpl{db: db}
}

func (rr *RecruitmentRepositoryImpl) CreateRecruitment(recruitment entity.Recruitment) error {
	result := rr.db.Create(&recruitment)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result.Error
}

func (rr *RecruitmentRepositoryImpl) GetRecruitments(limit int, offset int) ([]entity.Recruitment, error) {
	var recruitments []entity.Recruitment
	result := rr.db.Scopes(pagination.Paginate(limit, offset)).Preload(clause.Associations).Find(&recruitments)

	if result.Error != nil {
		return []entity.Recruitment{}, result.Error
	}

	return recruitments, nil
}

func (rr *RecruitmentRepositoryImpl) UpdateRecruitment(recruitment entity.Recruitment) error {
	result := rr.db.Model(&recruitment).Where("id = ?", recruitment.ID).Updates(recruitment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) CreateRecruitmentApplication(recruitmentApplication entity.RecruitmentApplication) error {
	result := rr.db.Create(&recruitmentApplication)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) GetRecruitmentByID(id uint) (entity.Recruitment, error) {
	var recruitment entity.Recruitment
	result := rr.db.Joins("Team").First(&recruitment, "recruitments.id = ?", id)
	if result.Error != nil {
		return recruitment, result.Error
	}

	return recruitment, nil
}

func (rr *RecruitmentRepositoryImpl) GetRecruitmentApplicationByRecruitmentID(id uint) ([]entity.RecruitmentApplication, error) {
	var recruitmentApplications []entity.RecruitmentApplication
	result := rr.db.Joins("User").Find(&recruitmentApplications, "recruitment_id = ?", id)
	if result.Error != nil {
		return recruitmentApplications, result.Error
	}

	return recruitmentApplications, nil
}

func (rr *RecruitmentRepositoryImpl) GetRecruitmentByUserID(id uint) ([]entity.Recruitment, error) {
	var recruitments []entity.Recruitment
	result := rr.db.Preload(clause.Associations).Find(&recruitments)
	if result.Error != nil {
		return recruitments, result.Error
	}

	return recruitments, nil
}

func (rr *RecruitmentRepositoryImpl) RejectRecruitmentApplication(id uint) error {
	var recruitmentApplication entity.RecruitmentApplication
	result := rr.db.First(&recruitmentApplication, id)
	if result.Error != nil {
		return result.Error
	}

	recruitmentApplication.IsAccepted = 0
	recruitmentApplication.IsRejected = 1
	result = rr.db.Save(recruitmentApplication)

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) AcceptRecruitmentApplication(id uint) error {
	var recruitmentApplication entity.RecruitmentApplication
	result := rr.db.First(&recruitmentApplication, id)
	if result.Error != nil {
		return result.Error
	}

	recruitmentApplication.IsAccepted = 1
	recruitmentApplication.IsRejected = 0
	result = rr.db.Save(recruitmentApplication)

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) GetRecruitmentApplicationByID(id uint) (entity.RecruitmentApplication, error) {
	var recruitmentApplication entity.RecruitmentApplication
	result := rr.db.Preload(clause.Associations).Find(&recruitmentApplication)
	fmt.Println(result)
	if result.Error != nil {
		return recruitmentApplication, result.Error
	}

	return recruitmentApplication, nil
}

func (rr *RecruitmentRepositoryImpl) DeleteRecruitmentByID(id uint) error {
	result := rr.db.Delete(&entity.Recruitment{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) OpenRecruitmentApplicationPeriod(id uint) error {
	result := rr.db.Model(&entity.Recruitment{}).Where("id = ?", id).Update("application_acceptance_status", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (rr *RecruitmentRepositoryImpl) CloseRecruitmentApplicationPeriod(id uint) error {
	result := rr.db.Model(&entity.Recruitment{}).Where("id = ?", id).Update("application_acceptance_status", 2)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}
