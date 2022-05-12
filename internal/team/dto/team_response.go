package dto

type TeamDetailsResponse struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Capacity    uint                 `json:"capacity"`
	TeamMembers []TeamMemberResponse `json:"members"`
}

type TeamMemberResponse struct {
	UserID   uint   `json:"id"`
	Name     string `json:"name"`
	IsLeader uint   `json:"isLeader"`
}
