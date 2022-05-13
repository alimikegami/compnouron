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
