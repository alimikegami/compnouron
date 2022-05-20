package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/team/entity"
)

type Recruitment struct {
	ID                          uint
	Role                        string      `gorm:"not null"`
	Description                 string      `gorm:"not null"`
	TeamID                      uint        `gorm:"not null"`
	Team                        entity.Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ApplicationAcceptanceStatus uint8       `gorm:"not null"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	RecruitmentApplications     []RecruitmentApplication
}
