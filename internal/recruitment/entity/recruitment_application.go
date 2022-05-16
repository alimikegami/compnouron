package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/user/entity"
)

type RecruitmentApplication struct {
	ID               uint
	UserID           uint
	RecruitmentID    uint
	AcceptanceStatus uint8
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Recruitment      Recruitment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User             entity.User
}
