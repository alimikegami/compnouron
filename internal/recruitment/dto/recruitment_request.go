package dto

type RecruitmentRequest struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	TeamID      uint   `json:"teamID"`
}

type RecruitmentApplicationRequest struct {
	RecruitmentID uint `json:"recruitmentID"`
	UserID        uint `json:"userID"`
}
