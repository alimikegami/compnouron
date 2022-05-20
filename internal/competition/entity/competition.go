package entity

import (
	"time"

	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type Competition struct {
	ID                       uint   `gorm:"primaryKey"`
	Name                     string `gorm:"not null"`
	Description              string `gorm:"not null"`
	ContactPerson            string `gorm:"not null"`
	IsTeam                   int8   `gorm:"not null"`
	IsTheSameInstitution     int8   `gorm:"not null"`
	RegistrationPeriodStatus int8   `gorm:"not null"`
	TeamCapacity             int8   `gorm:"not null"`
	Level                    string `gorm:"not null"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
	UserID                   uint `gorm:"not null"`
	User                     userEntity.User
	CompetitionRegistrations []CompetitionRegistration
}
