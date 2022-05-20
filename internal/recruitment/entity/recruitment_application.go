package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/user/entity"
)

type RecruitmentApplication struct {
	ID               uint
	UserID           uint  `gorm:"not null"`
	RecruitmentID    uint  `gorm:"not null"`
	AcceptanceStatus uint8 `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Recruitment      Recruitment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User             entity.User
}
