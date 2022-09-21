package model

type Notification struct {
	ID         string `json:"id"`
	FromUser   *User  `json:"fromUser"`
	FromUserID string `json:"fromUserId"`
	ToUser     *User  `json:"toUser"`
	ToUserID   string `json:"toUserId"`
	Message    string `json:"text"`
}
