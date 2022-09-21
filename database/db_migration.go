package database

import "github.com/renaldiaddison/tpa-web-backend/graph/model"

func MigrateDatabase() {
	db := GetDatabase()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ActivationLink{})
	db.AutoMigrate(&model.ResetPasswordLink{})
	db.AutoMigrate(&model.Education{})
	db.AutoMigrate(&model.Experience{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Hashtag{})
	db.AutoMigrate(&model.Job{})
	db.AutoMigrate(&model.LikeComment{})
}
