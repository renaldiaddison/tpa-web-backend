package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func GetDatabase() *gorm.DB {

	if database == nil {
		dsn := "host=localhost user=postgres password=postgres dbname=linkHEdin port=5432 sslmode=disable TimeZone=Asia/Bangkok"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			database = db
		} else {
			panic(err)
		}
	}

	return database
}
