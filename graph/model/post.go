package model

import "time"

type Post struct {
	ID        string     `json:"ID"`
	Text      string     `json:"text"`
	URL       string     `json:"url"`
	PhotoUrl  string     `json:"photoUrl"`
	VideoUrl  string     `json:"videoUrl"`
	SenderId  string     `json:"senderId"`
	CreatedAt time.Time  `json:"createdAt"`
	Sender    *User      `json:"Sender" gorm:"reference:User"`
	Likes     []*User    `json:"Likes" gorm:"many2many:like_posts"`
	Comment   []*Comment `json:"Comment" gorm:"foreignKey:PostID;"`
}

type LikePosts struct {
	PostId string `json:"PostId"`
	UserId string `json:"UserId"`
}

type Comment struct {
	ID               string         `json:"id"`
	Comment          string         `json:"comment"`
	PostID           string         `json:"postId"`
	CommenterID      string         `json:"commenterId"`
	Commenter        *User          `json:"Commenter"`
	ReplyToCommentID *string        `json:"replyToCommentId"`
	LikeComment      []*LikeComment `json:"LikeComment" gorm:"foreignKey:CommentID"`
	Replies          []*Comment     `json:"Replies" gorm:"foreignKey:ReplyToCommentID"`
	CreatedAt        time.Time      `json:"createdAt"`
}

type LikeComment struct {
	ID        string `json:"id"`
	CommentID string `json:"commentID"`
	UserID    string `json:"userID"`
	User      *User  `json:"User"`
}
