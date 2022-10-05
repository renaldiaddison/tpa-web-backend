package model

import "time"

type Room struct {
	ID        string `json:"id"`
	User1ID   string
	User1     *User `json:"user1"`
	User2ID   string
	User2     *User      `json:"user2"`
	Messages  []*Message `json:"messages"`
	CreatedAt time.Time  `json:"createdAt"`
}

type Message struct {
	ID             string `json:"id"`
	Sender         *User  `json:"sender"`
	SharePost      *Post  `json:"SharePost"`
	ShareProfile   *User  `json:"ShareProfile"`
	Text           string `json:"text"`
	ImageURL       string `json:"imageUrl"`
	SenderID       string
	RoomID         string
	ShareProfileID *string
	SharePostID    *string
	CreatedAt      time.Time `json:"createdAt"`
}
