package entity

import userEntity "github.com/alimikegami/compnouron/internal/user/entity"

type TeamMember struct {
	ID       uint `gorm:"primaryKey"`
	TeamID   uint
	UserID   uint
	IsLeader uint
	User     userEntity.User
	Team     Team
}
