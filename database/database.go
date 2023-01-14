package database

import (
	"fmt"
	"log"
	"os"

	"github.com/meme_hub/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connectDB() {

	// Loading environment variables
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// Database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	// Opening connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	
	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Make migrations to the database if they have not already been created
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})


	Database = DbInstance{Db:db}
}