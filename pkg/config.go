package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	// สร้าง URL สำหรับการเชื่อมต่อ MySQL
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	return &Config{
		DatabaseURL: databaseURL,
	}
}
