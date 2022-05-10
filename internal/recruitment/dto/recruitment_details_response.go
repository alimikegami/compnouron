package dto

type RecruitmentDetailsResponse struct {
	Recruitment             RecruitmentResponse              `json:"recruitment"`
	RecruitmentApplications []RecruitmentApplicationResponse `json:"recruitmentApplications"`
}
