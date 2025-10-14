package config

import (
	"fullstack/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Car{})
	if err != nil {
		panic(err)
	}

	DB = db
	log.Println("Success connect Postgres")
}
