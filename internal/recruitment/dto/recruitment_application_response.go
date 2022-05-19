package dto

type RecruitmentApplicationResponse struct {
	ID               uint   `json:"ID"`
	UserID           uint   `json:"userID"`
	UserName         string `json:"userName"`
	RecruitmentID    uint   `json:"recruitmentID"`
	AcceptanceStatus uint8  `json:"acceptanceStatus"`
}
