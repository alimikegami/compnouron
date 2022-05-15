package repository

import (
	"github.com/alimikegami/compnouron/db/pagination"

	"errors"

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

func (cr *CompetitionRepository) GetCompetitions(limit int, offset int) ([]entity.Competition, error) {
	var competitions []entity.Competition
	result := cr.db.Scopes(pagination.Paginate(limit, offset)).Find(&competitions)

	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return competitions, nil
}

func (cr *CompetitionRepository) Register(competitionRegistration entity.CompetitionRegistration) error {
	result := cr.db.Create(&competitionRegistration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompetitionRepository) GetCompetitionRegistration(competitionID uint) (entity.Competition, error) {
	var competitionRegistration entity.Competition

	result := cr.db.Preload("CompetitionRegistrations", "acceptance_status = 0").Where("id = ?", competitionID).Find(&competitionRegistration)
	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competitionRegistration, nil
}

func (cr *CompetitionRepository) GetAcceptedCompetitionParticipants(competitionID uint) (entity.Competition, error) {
	var competition entity.Competition

	result := cr.db.Preload("CompetitionRegistrations", "acceptance_status = 1").Where("id = ?", competitionID).Find(&competition)

	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competition, nil
}

func (cr *CompetitionRepository) RejectCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("acceptance_status", 2)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepository) AcceptCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("acceptance_status", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepository) CloseCompetitionRegistrationPeriod(id uint) error {
	result := cr.db.Model(&entity.Competition{}).Where("id = ?", id).Update("is_open", 2)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepository) OpenCompetitionRegistrationPeriod(id uint) error {
	result := cr.db.Model(&entity.Competition{}).Where("id = ?", id).Update("registration_period_status", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepository) SearchCompetition(limit int, offset int, keyword string) ([]entity.Competition, error) {
	var competitions []entity.Competition
	result := cr.db.Scopes(pagination.Paginate(limit, offset)).Where("name LIKE ?", "%"+keyword+"%").Find(&competitions)
	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return competitions, nil
}
