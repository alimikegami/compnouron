package repository

import (
	"errors"

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
	result := cr.db.First(&competition, ID)

	if result.Error != nil {
		return entity.Competition{}, nil
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

func (cr *CompetitionRepository) Register(competitionRegistration entity.CompetitionRegistration) error {
	result := cr.db.Create(&competitionRegistration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompetitionRepository) GetCompetitionRegistration(competitionID uint) ([]entity.CompetitionRegistration, error) {
	var competitionRegistration []entity.CompetitionRegistration

	result := cr.db.Preload(clause.Associations).Where("competition_id = ?", competitionID).Find(&competitionRegistration)
	if result.Error != nil {
		return []entity.CompetitionRegistration{}, result.Error
	}

	return competitionRegistration, nil
}

func (cr *CompetitionRepository) RejectCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("is_accepted", 0)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepository) AcceptCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("is_accepted", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}
