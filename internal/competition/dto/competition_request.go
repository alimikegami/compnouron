package dto

type CompetitionRequest struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	ContactPerson        string `json:"contactPerson"`
	IsTheSameInstitution int8   `json:"isTheSameInstitution"`
	IsTeam               int8   `json:"isTeam"`
	TeamCapacity         int8   `json:"teamCapacity"`
	Level                string `json:"level"`
}
