package entity

import (
	"time"
)

type User struct {
	ID                uint `gorm:"primaryKey"`
	Name              string
	Email             string `gorm:"unique"`
	PhoneNumber       string
	Password          string
	SchoolInstitution string
	Skills            []Skill
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
