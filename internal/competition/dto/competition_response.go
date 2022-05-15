package dto

type CompetitionResponse struct {
	ID            uint   `json:"ID"`
	Name          string `json:"name"`
	ContactPerson string `json:"contactPerson"`
	IsTeam        int8   `json:"isTeam"`
	Level         string `json:"level"`
}
