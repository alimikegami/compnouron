package repository

import (
	"fmt"

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
