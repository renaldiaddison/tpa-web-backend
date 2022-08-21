package model

type ResetPasswordLink struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
}
