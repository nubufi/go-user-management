package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"goUserManagement/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDb connects to the postgresql database
func ConnectToDb() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if host == "" {
		log.Fatal(errors.New("DB_HOST environment variable is not set"))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, port, dbName)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		time.Sleep(5 * time.Second)
		ConnectToDb()
	}

	log.Println("Connected to database")

	return DB
}

// Migrate migrates the models
func Migrate(db *gorm.DB) {
	// Migrate migrates the models
	db.AutoMigrate(
		&models.User{},
	)
}
