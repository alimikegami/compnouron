package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/team/entity"
)

type Recruitment struct {
	ID          uint        `json:"id"`
	Role        string      `json:"role"`
	Description string      `json:"description"`
	TeamID      uint        `json:"teamID"`
	Team        entity.Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
