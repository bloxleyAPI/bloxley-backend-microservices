package main

import (
	"auth-microservice/config"
	"auth-microservice/models"
	"auth-microservice/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration and connect to the database
	config.Load()

	// Run migrations
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully.")

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
