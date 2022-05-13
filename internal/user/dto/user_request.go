package dto

type SkillRequest struct {
	Name string `json:"name"`
}

type UserRegistrationRequest struct {
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	PhoneNumber       string         `json:"phoneNumber"`
	Password          string         `json:"password"`
	SchoolInstitution string         `json:"schoolInstitution"`
	Skills            []SkillRequest `json:"skills"`
}
