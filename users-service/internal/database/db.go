package database

import (
	"fmt"
	"log"
	"os"

	"github.com/PBGlory/users-service/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "users_service")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, dbUser, password, dbname,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, _ := DB.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Cannot ping PostgreSQL: %v", err)
	}

	// Автомиграция
	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Connected to PostgreSQL and migrated successfully")
}

func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
