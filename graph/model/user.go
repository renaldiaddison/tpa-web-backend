package model

type User struct {
	ID                string            `json:"id" gorm:"primaryKey"`
	Email             string            `json:"email"`
	FirstName         string            `json:"firstName"`
	LastName          string            `json:"lastName"`
	AdditionalName    string            `json:"additionalName"`
	Password          string            `json:"password"`
	IsActive          bool              `json:"is_active"`
	ProfilePicture    string            `json:"profile_picture"`
	BackgroundPicture string            `json:"background_picture"`
	Headline          string            `json:"headline"`
	About             string            `json:"about"`
	Location          string            `json:"location"`
	ProfileViews      int               `json:"profile_views"`
	Visits            []*User           `json:"Visit" gorm:"many2many:user_visits"`
	Follows           []*User           `json:"Follow" gorm:"many2many:user_follows"`
	Comment           []*Comment        `json:"Comment" gorm:"foreignKey:CommenterID;"`
	LikeComment       []*LikeComment    `json:"LikeComment" gorm:"foreignKey:UserID"`
	Connection        []*Connection     `json:"Connection" gorm:"foreignKey:User1ID;foreignKey:User2ID"`
	ConnectRequest    []*ConnectRequest `json:"ConnectRequest" gorm:"foreignKey:FromUserID;foreignKey:ToUserID"`
	Block             []*User           `json:"Block" gorm:"many2many:user_blocks"`
	Experiences       []*Experience     `json:"Experiences" gorm:"foreignKey:UserID"`
	Educations        []*Education      `json:"Educations" gorm:"foreignKey:UserID"`
	Notification      []*Notification   `json:"Notification" gorm:"foreignKey:FromUserID;foreignKey:ToUserID"`
}

type Connection struct {
	ID      string `json:"id"`
	User1   *User  `json:"user1"`
	User1ID string `json:"user1Id"`
	User2   *User  `json:"user2"`
	User2ID string `json:"user2Id"`
}

type Visit struct {
	UserID  string `json:"userId"`
	VisitID string `json:"visitId"`
}

type Follow struct {
	UserID   string `json:"userId"`
	FollowID string `json:"followId"`
}

type ConnectRequest struct {
	ID         string `json:"id"`
	FromUserID string `json:"fromUserId"`
	FromUser   *User  `json:"fromUser"`
	ToUser     *User  `json:"toUser"`
	ToUserID   string `json:"toUserId"`
	Message    string `json:"message"`
}
