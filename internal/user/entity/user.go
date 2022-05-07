package entity

import "time"

type User struct {
	ID                uint `gorm:"primaryKey"`
	Name              string
	Email             string
	PhoneNumber       string
	Password          string
	SchoolInstitution string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Skill struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Skills []Skill
