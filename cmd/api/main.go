package main

import (
	"carSearch/config"
	"carSearch/frontend"
	"carSearch/internal/handler"
	"carSearch/internal/repository"
	"carSearch/internal/service"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	carRepo := repository.NewCarRepository(db)
	producerRepo := repository.NewProducerRepository(db)

	// Initialize services
	carService := service.NewCarService(carRepo, producerRepo)

	// Initialize handlers
	carHandler := handler.NewCarHandler(carService, cfg.APIKey)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	})

	// Add CORS middleware to allow frontend access
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://localhost:3000,https://car.marc-schulz.online", // Common Vite/SvelteKit ports
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-API-Key",
		AllowCredentials: true,
	}))

	// Register routes
	handler.RegisterRoutes(app, carHandler)
	app.Use("/*", filesystem.New(filesystem.Config{
		Root:         frontend.Dist(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))
	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
