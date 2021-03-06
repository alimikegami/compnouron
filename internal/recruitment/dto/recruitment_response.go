package dto

type RecruitmentResponse struct {
	ID                          uint   `json:"ID"`
	Role                        string `json:"role"`
	Description                 string `json:"description"`
	TeamID                      uint   `json:"teamID"`
	TeamName                    string `json:"teamName"`
	ApplicationAcceptanceStatus uint8  `json:"ApplicationAcceptanceStatus"`
}

type BriefRecruitmentResponse struct {
	ID                          uint   `json:"ID"`
	Role                        string `json:"role"`
	TeamName                    string `json:"teamName"`
	ApplicationAcceptanceStatus uint8  `json:"ApplicationAcceptanceStatus"`
}

type RecruitmentsResponse []RecruitmentResponse
