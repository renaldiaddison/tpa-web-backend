package model

type Experience struct {
	ID             string `json:"ID" gorm:"primaryKey"`
	UserID         string `json:"UserID"`
	Title          string `json:"Title"`
	EmploymentType string `json:"EmploymentType"`
	CompanyName    string `json:"CompanyName"`
	Location       string `json:"Location"`
	Active         bool   `json:"Active"`
	StartYear      string `json:"StartYear"`
	EndYear        string `json:"EndYear"`
	Industry       string `json:"Industry"`
	Description    string `json:"Description"`
}
