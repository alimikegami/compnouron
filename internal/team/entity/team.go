package entity

import "time"

type Team struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Capacity    uint   `gorm:"not null"`
	TeamMembers []TeamMember
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
