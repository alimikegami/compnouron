package entity

import (
	"time"

	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type TeamMember struct {
	ID        uint `gorm:"primaryKey"`
	TeamID    uint
	UserID    uint
	IsLeader  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userEntity.User
	Team      Team
}
