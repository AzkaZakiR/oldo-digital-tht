package database

import (
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func OpenConnection() (*gorm.DB, error) {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "oldo_db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_TIMEZONE", "UTC"),
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to database")
			return db, nil
		}
		log.Println("⏳ Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect after retries: %w", err)
}
func Migrate(db *gorm.DB) error{
	err := db.AutoMigrate(
		&models.User{},
		&models.DataPlan{},
		&models.Transaction{},
	)
	if err != nil{
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migration successfull")
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}