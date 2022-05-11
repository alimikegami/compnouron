package dto

import "time"

type DetailedCompetitionResponse struct {
	ID                   uint      `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	ContactPerson        string    `json:"contactPerson"`
	IsTheSameInstitution int8      `json:"isTheSameInstitution"`
	IsTeam               int8      `json:"isTeam"`
	TeamCapacity         int8      `json:"teamCapacity"`
	Level                string    `json:"level"`
	UpdatedAt            time.Time `json:"updatedAt"`
	UserID               uint      `json:"userID"`
	UserName             string    `json:"userName"`
}
