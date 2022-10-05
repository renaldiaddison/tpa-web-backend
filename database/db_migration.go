package database

import "github.com/renaldiaddison/tpa-web-backend/graph/model"

func MigrateDatabase() {
	db := GetDatabase()
	db.AutoMigrate(&model.User{},
		&model.ActivationLink{},
		&model.ResetPasswordLink{},
		&model.Education{},
		&model.Experience{},
		&model.Post{},
		&model.Block{},
		&model.ConnectRequest{},
		&model.Connection{},
		&model.Comment{},
		&model.Hashtag{},
		&model.Job{},
		&model.LikeComment{},
		&model.Notification{},
		&model.Room{},
		&model.Message{},
	)
}
