package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/team/entity"
	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type CompetitionRegistration struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint `gorm:"default:null"`
	TeamID           uint `gorm:"default:null"`
	CompetitionID    uint
	AcceptanceStatus uint `gorm:"default:null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Team             entity.Team
	User             userEntity.User
	Competition      Competition
}
