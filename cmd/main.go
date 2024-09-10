package main

import (
	"github.com/be2bag/myapp/internal/core/services"
	"github.com/be2bag/myapp/internal/handlers"
	"github.com/be2bag/myapp/internal/repositories"
	"github.com/be2bag/myapp/pkg"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Load configuration
	config := pkg.LoadConfig()

	// Initialize database connection
	db, err := gorm.Open(mysql.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	handlers.NewUserHandler(app, userService)

	// Start the server
	app.Listen(":3000")
}
