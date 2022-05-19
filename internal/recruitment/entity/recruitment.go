package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/team/entity"
)

type Recruitment struct {
	ID                          uint
	Role                        string
	Description                 string
	TeamID                      uint
	Team                        entity.Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ApplicationAcceptanceStatus uint8
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	RecruitmentApplications     []RecruitmentApplication
}
