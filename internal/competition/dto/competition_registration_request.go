package dto

type CompetitionRegistrationRequest struct {
	UserID        uint `json:"userID"`
	TeamID        uint `json:"teamID"`
	CompetitionID uint `json:"competitionID"`
}
