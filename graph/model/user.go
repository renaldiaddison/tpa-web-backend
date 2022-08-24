package model

import (
	"github.com/lib/pq"
)

type User struct {
	ID                string         `json:"id" gorm:"primaryKey"`
	Email             string         `json:"email"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	AdditionalName    string         `json:"additionalName"`
	Password          string         `json:"password"`
	IsActive          bool           `json:"is_active"`
	ProfilePicture    string         `json:"profile_picture"`
	BackgroundPicture string         `json:"background_picture"`
	Headline          string         `json:"headline"`
	About             string         `json:"about"`
	Location          string         `json:"location"`
	ProfileViews      int            `json:"profile_views"`
	FollowedUser      pq.StringArray `json:"followed_user" gorm:"type:text[]"`
	RequestConnect    pq.StringArray `json:"request_connect" gorm:"type:text[]"`
	ConnectedUser     pq.StringArray `json:"connected_user" gorm:"type:text[]"`
}
