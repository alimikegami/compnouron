package repository

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"gorm.io/gorm"
)

type RecruitmentRepository struct {
	db *gorm.DB
}

func CreateNewRecruitmentRepository(db *gorm.DB) *RecruitmentRepository {
	return &RecruitmentRepository{db: db}
}

func (rr *RecruitmentRepository) CreateRecruitment(recruitment entity.Recruitment) error {
	result := rr.db.Create(&recruitment)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result.Error
}
