package dto

type SkillResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"userID"`
}

type UserDetailsResponse struct {
	ID                uint            `json:"id"`
	Name              string          `json:"name"`
	Email             string          `json:"email"`
	PhoneNumber       string          `json:"phoneNumber"`
	SchoolInstitution string          `json:"schoolInstitution"`
	Skills            []SkillResponse `json:"skills"`
}

type UserCompetitionHistory struct {
	CompetitionRegistrationID uint   `json:"id"`
	CompetitionID             uint   `json:"competitionID"`
	CompetitionName           string `json:"competitionName"`
	AcceptanceStatus          uint   `json:"acceptanceStatus"`
}

type UserRecruitmentApplicationHistory struct {
	RecruitmentApplicationID uint   `json:"id"`
	RecruitmentID            uint   `json:"recruitmentID"`
	RecruitmentRole          string `json:"recruitmentRole"`
	AcceptanceStatus         uint   `json:"acceptanceStatus"`
}
