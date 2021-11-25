package repository

import (
	"golang-gin/entities"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + database + " sslmode=require password=" + password

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entities.Video{})

	DB = db
}
