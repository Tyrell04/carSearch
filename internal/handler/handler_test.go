package handler

import (
	"carSearch/config"
	"carSearch/internal/domain"
	"carSearch/internal/repository"
	"carSearch/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestServer creates a configured Fiber app for testing
func setupTestServer() *fiber.App {
	// Initialize in-memory database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Car{}, &domain.Producer{})

	// Initialize repositories
	carRepo := repository.NewCarRepository(db)
	producerRepo := repository.NewProducerRepository(db)

	// Initialize service
	carService := service.NewCarService(carRepo, producerRepo)

	// Load test configuration (API Key)
	cfg := &config.Config{
		APIKey: "test-api-key", // Use a dummy API key for testing
	}

	// Initialize handler
	carHandler := NewCarHandler(carService, cfg.APIKey)

	// Initialize Fiber app
	app := fiber.New()

	// Define test routes
	RegisterRoutes(app, carHandler)

	return app
}
