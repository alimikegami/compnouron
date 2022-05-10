package dto

type RecruitmentRequest struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	TeamID      uint   `json:"teamID"`
}
