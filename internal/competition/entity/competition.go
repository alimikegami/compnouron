package entity

import (
	"time"

	userEntity "github.com/alimikegami/compnouron/internal/user/entity"
)

type Competition struct {
	ID                   uint `gorm:"primaryKey"`
	Name                 string
	Description          string
	ContactPerson        string
	IsTheSameInstitution int8
	IsTeam               int8
	TeamCapacity         int8
	Level                string
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
	UserID               uint
	User                 userEntity.User
}
