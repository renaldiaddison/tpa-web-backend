package model

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
}
