package database

import (
	"go-api/src/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(" -> Could not connect to the database")
	}

	DB = database
	_ = database.AutoMigrate(&models.User{})
}
