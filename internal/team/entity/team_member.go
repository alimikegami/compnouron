package entity

import userEntity "github.com/alimikegami/compnouron/internal/user/entity"

type TeamMember struct {
	ID     uint `json:"id"`
	TeamID uint `json:"teamId"`
	UserID uint `json:"userId"`
	User   userEntity.User
	Team   Team
}