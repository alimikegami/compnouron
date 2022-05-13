package dto

type TeamCompetitionRegistrationResponse struct {
	ID            uint
	TeamID        uint
	TeamName      string
	CompetitionID uint
	IsAccepted    uint
}
