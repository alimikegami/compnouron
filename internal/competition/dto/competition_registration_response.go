package dto

type TeamCompetitionRegistrationResponse struct {
	ID            uint   `json:"id"`
	TeamID        uint   `json:"teamID"`
	TeamName      string `json:"teamName"`
	CompetitionID uint   `json:"competitionID"`
	IsAccepted    uint   `json:"isAccepted"`
}

type IndividualCompetitionRegistrationResponse struct {
	ID                uint   `json:"id"`
	UserID            uint   `json:"userID"`
	UserName          string `json:"userName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	SchoolInstitution string `json:"schoolInstitution"`
	CompetitionID     uint   `json:"competitionID"`
	IsAccepted        uint   `json:"isAccepted"`
}
