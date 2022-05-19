package repository

import (
	"github.com/alimikegami/compnouron/db/pagination"

	"errors"

	"github.com/alimikegami/compnouron/internal/competition/entity"
	"gorm.io/gorm"
)

type CompetitionRepositoryImpl struct {
	db *gorm.DB
}

type CompetitionRepository interface {
	CreateCompetition(competition *entity.Competition) error
	DeleteCompetition(ID uint) error
	GetCompetitionByID(ID uint) (entity.Competition, error)
	GetCompetitionByUserID(userID uint) ([]entity.Competition, error)
	UpdateCompetition(competition entity.Competition) error
	GetCompetitions(limit int, offset int) ([]entity.Competition, error)
	Register(competitionRegistration entity.CompetitionRegistration) error
	GetCompetitionRegistration(competitionID uint) (entity.Competition, error)
	GetCompetitionRegistrationByUserID(userID uint) ([]entity.CompetitionRegistration, error)
	GetAcceptedCompetitionParticipants(competitionID uint) (entity.Competition, error)
	RejectCompetitionRegistration(id uint) error
	AcceptCompetitionRegistration(id uint) error
	CloseCompetitionRegistrationPeriod(id uint) error
	OpenCompetitionRegistrationPeriod(id uint) error
	SearchCompetition(limit int, offset int, keyword string) ([]entity.Competition, error)
}

func CreateNewCompetitionRepository(db *gorm.DB) CompetitionRepository {
	return &CompetitionRepositoryImpl{db: db}
}

func (cr *CompetitionRepositoryImpl) CreateCompetition(competition *entity.Competition) error {
	result := cr.db.Create(&competition)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) GetCompetitionRegistrationByUserID(userID uint) ([]entity.CompetitionRegistration, error) {
	var compRegistration []entity.CompetitionRegistration
	result := cr.db.Joins("Competition").Find(&compRegistration, "competition_registrations.user_id = ?", userID)
	if result.Error != nil {
		return []entity.CompetitionRegistration{}, result.Error
	}
	return compRegistration, nil
}

func (cr *CompetitionRepositoryImpl) DeleteCompetition(ID uint) error {
	result := cr.db.Delete(&entity.Competition{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompetitionRepositoryImpl) GetCompetitionByID(ID uint) (entity.Competition, error) {
	var competition entity.Competition
	result := cr.db.First(&competition, ID)

	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competition, nil
}

func (cr *CompetitionRepositoryImpl) UpdateCompetition(competition entity.Competition) error {
	result := cr.db.Model(&competition).Where("id = ?", competition.ID).Updates(competition)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) GetCompetitions(limit int, offset int) ([]entity.Competition, error) {
	var competitions []entity.Competition
	result := cr.db.Scopes(pagination.Paginate(limit, offset)).Find(&competitions)

	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return competitions, nil
}

func (cr *CompetitionRepositoryImpl) Register(competitionRegistration entity.CompetitionRegistration) error {
	result := cr.db.Create(&competitionRegistration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompetitionRepositoryImpl) GetCompetitionRegistration(competitionID uint) (entity.Competition, error) {
	var competitionRegistration entity.Competition

	result := cr.db.Preload("CompetitionRegistrations", "acceptance_status = 0").Where("id = ?", competitionID).Find(&competitionRegistration)
	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competitionRegistration, nil
}

func (cr *CompetitionRepositoryImpl) GetAcceptedCompetitionParticipants(competitionID uint) (entity.Competition, error) {
	var competition entity.Competition

	result := cr.db.Preload("CompetitionRegistrations", "acceptance_status = 1").Where("id = ?", competitionID).Find(&competition)

	if result.Error != nil {
		return entity.Competition{}, result.Error
	}

	return competition, nil
}

func (cr *CompetitionRepositoryImpl) RejectCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("acceptance_status", 2)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) AcceptCompetitionRegistration(id uint) error {
	result := cr.db.Model(&entity.CompetitionRegistration{}).Where("id = ?", id).Update("acceptance_status", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) CloseCompetitionRegistrationPeriod(id uint) error {
	result := cr.db.Model(&entity.Competition{}).Where("id = ?", id).Update("is_open", 2)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) OpenCompetitionRegistrationPeriod(id uint) error {
	result := cr.db.Model(&entity.Competition{}).Where("id = ?", id).Update("registration_period_status", 1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (cr *CompetitionRepositoryImpl) GetCompetitionByUserID(userID uint) ([]entity.Competition, error) {
	var comps []entity.Competition
	result := cr.db.Find(comps, "user_id = ?", userID)
	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return comps, nil
}

func (cr *CompetitionRepositoryImpl) SearchCompetition(limit int, offset int, keyword string) ([]entity.Competition, error) {
	var competitions []entity.Competition
	result := cr.db.Scopes(pagination.Paginate(limit, offset)).Where("name LIKE ?", "%"+keyword+"%").Find(&competitions)
	if result.Error != nil {
		return []entity.Competition{}, result.Error
	}

	return competitions, nil
}
