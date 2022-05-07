package dto

type TeamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    uint   `json:"capacity"`
}
