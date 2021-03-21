package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
