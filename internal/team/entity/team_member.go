package entity

import (
	"time"

	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type TeamMember struct {
	ID        uint `gorm:"primaryKey"`
	TeamID    uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	IsLeader  uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userEntity.User
	Team      Team
}
