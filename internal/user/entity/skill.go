package entity

import "time"

type Skill struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
