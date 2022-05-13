package entity

import "time"

type Skill struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
