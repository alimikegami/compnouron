package entity

import (
	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type Team struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    uint   `json:"capacity"`
	UserID      uint   `json:"userID"`
	User        userEntity.User
}
