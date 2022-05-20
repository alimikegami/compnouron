package entity

import (
	"time"
)

type User struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Email             string `gorm:"unique;not null"`
	PhoneNumber       string `gorm:"not null"`
	Password          string `gorm:"not null"`
	SchoolInstitution string `gorm:"not null"`
	Skills            []Skill
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
