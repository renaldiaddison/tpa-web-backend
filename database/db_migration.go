package database

import "github.com/renaldiaddison/tpa-web-backend/graph/model"

func MigrateDatabase() {
	db := GetDatabase()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ActivationLink{})
	db.AutoMigrate(&model.ResetPasswordLink{})
	db.AutoMigrate(&model.Education{})
	db.AutoMigrate(&model.Experience{})
}
