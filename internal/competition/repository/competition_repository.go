package repository

import (
	"github.com/alimikegami/compnouron/internal/competition/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (cr *CompetitionRepository) GetCompetitionByID(ID uint) (entity.Competition, error) {
	var competition entity.Competition
	result := cr.db.Preload(clause.Associations).First(&competition, ID)
	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competition, nil
}

func (cr *CompetitionRepository) UpdateCompetition(competition entity.Competition) error {
	result := cr.db.Model(&competition).Where("id = ?", competition.ID).Updates(competition)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
