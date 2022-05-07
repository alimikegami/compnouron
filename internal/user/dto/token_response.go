package dto

type TokenResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
}
