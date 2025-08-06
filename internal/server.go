package internal

import (
	"carSearch/config"
	"carSearch/frontend"
	"carSearch/internal/domain"
	"carSearch/internal/handler"
	"carSearch/internal/repository"
	"carSearch/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Start(cfg *config.Config) {
	// Initialize database
	gormDB, err := gorm.Open(sqlite.Open("carsearch.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate database schema
	gormDB.AutoMigrate(&domain.Car{}, &domain.Producer{})

	// Initialize repositories
	carRepo := repository.NewCarRepository(gormDB)
	producerRepo := repository.NewProducerRepository(gormDB)

	// Initialize service and handler
	carService := service.NewCarService(carRepo, producerRepo)
	carHandler := handler.NewCarHandler(carService, cfg.APIKey)

	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	handler.RegisterRoutes(app, carHandler)

	app.Use("/*", filesystem.New(filesystem.Config{
		Root:         frontend.Dist(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))
	// Start server
	log.Fatal(app.Listen(":3000"))
}
