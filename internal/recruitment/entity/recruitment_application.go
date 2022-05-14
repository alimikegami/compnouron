package entity

import "github.com/alimikegami/compnouron/internal/user/entity"

type RecruitmentApplication struct {
	ID            uint
	UserID        uint
	RecruitmentID uint
	IsAccepted    uint8
	IsRejected    uint8
	Recruitment   Recruitment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User          entity.User
}
