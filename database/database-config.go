package database

import (
	"log"

	"github.com/jamesthomasw/go-mygram-v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	
	dsn := "postgres://postgres:kikohuang@localhost:5432/go-mygram"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	return db
}