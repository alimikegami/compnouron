package repository

import (
	"github.com/alimikegami/compnouron/db/pagination"
	"github.com/alimikegami/compnouron/internal/competition/entity"
	"gorm.io/gorm"
)

type CompetitionRepository struct {
	db *gorm.DB
}

func CreateNewCompetitionRepository(db *gorm.DB) *CompetitionRepository {
	return &CompetitionRepository{db: db}
}

func (cr *CompetitionRepository) CreateCompetition(competition *entity.Competition) error {
	result := cr.db.Create(&competition)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *CompetitionRepository) DeleteCompetition(ID uint) error {
	result := cr.db.Delete(&entity.Competition{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompetitionRepository) GetCompetitionByID(ID uint) *entity.Competition {
	var competition entity.Competition
	cr.db.First(&competition, ID)

	return &competition
}

func (cr *CompetitionRepository) UpdateCompetition(competition entity.Competition) error {
	result := cr.db.Model(&competition).Where("id = ?", competition.ID).Updates(competition)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *CompetitionRepository) GetCompetitions(limit int, offset int) ([]entity.Competition, error) {
	var competitions []entity.Competition
	result := cr.db.Scopes(pagination.Paginate(limit, offset)).Find(&competitions)

	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return competitions, nil
}
