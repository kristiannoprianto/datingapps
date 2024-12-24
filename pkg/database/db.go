package database

import (
	"apps/datingapps/pkg/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect connects to the database
func Connect() error {
	dsn := "host=localhost user=gorm password=gorm dbname=dating_app port=5432 sslmode=disable"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return err
	}

	DB = db
	return db.AutoMigrate(&models.User{}, &models.Swipe{})
}
