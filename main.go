package main

import (
	"log"
	"sample/api/routes"
	"sample/db"

	models "sample/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Auto Migrate
	db.DB.AutoMigrate(&models.User{})

	// Set up Fiber app
	app := fiber.New()

	// Register user routes
	routes.SetupUserRoutes(app)

	// Start server
	err = app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
