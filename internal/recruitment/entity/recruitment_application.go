package entity

import (
	"time"

	"github.com/alimikegami/compnouron/internal/user/entity"
)

type RecruitmentApplication struct {
	ID            uint
	UserID        uint
	RecruitmentID uint
	IsAccepted    uint8
	IsRejected    uint8
	IsOpen        uint8
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          entity.User
}
