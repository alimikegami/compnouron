package dto

type CompetitionResponse struct {
	ID            uint   `json:"ID"`
	Name          string `json:"name"`
	ContactPerson string `json:"contactPerson"`
	IsTeam        int8   `json:"isTeam"`
	Level         string `json:"level"`
}

type DetailedCompetitionResponse struct {
	ID                       uint `gorm:"primaryKey"`
	Name                     string
	Description              string
	ContactPerson            string
	IsTeam                   int8
	RegistrationPeriodStatus int8
	TeamCapacity             int8
	Level                    string
	UserID                   uint
	UserName                 string
}
