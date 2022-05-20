package dto

type BriefTeamResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TeamDetailsResponse struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Capacity    uint                 `json:"capacity"`
	TeamMembers []TeamMemberResponse `json:"members"`
}

type TeamMemberResponse struct {
	UserID            uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
	SchoolInstitution string `json:"schoolInstitution"`
	IsLeader          uint   `json:"isLeader"`
}
