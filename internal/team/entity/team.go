package entity

import "time"

type Team struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Capacity    uint
	TeamMembers []TeamMember
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
