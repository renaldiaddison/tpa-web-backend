package database

import "github.com/renaldiaddison/tpa-web-backend/graph/model"

func MigrateDatabase() {
	db := GetDatabase()
	db.AutoMigrate(&model.User{})
}
