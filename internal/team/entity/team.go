package entity

type Team struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Capacity    uint
}
