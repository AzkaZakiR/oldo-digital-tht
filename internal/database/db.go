package database

import (
	"fmt"
	"log"
	"os"

	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func OpenConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	dsn := "host=localhost user=postgres password=" + os.Getenv("DB_PASSWORD") + " dbname=oldo_digital port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error{
	err := db.AutoMigrate(
		&models.User{},
		&models.DataPlan{},
		&models.Transaction{},
	)
	if err != nil{
		log.Fatal("Migration failed", err)
	}

	log.Println("Migration successfull")
	return nil
}