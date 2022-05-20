package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/team/entity"
	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type CompetitionRegistration struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	TeamID           uint
	CompetitionID    uint `gorm:"not null"`
	AcceptanceStatus uint `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Team             entity.Team
	User             userEntity.User
	Competition      Competition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
