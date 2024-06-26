package main

import (
	"carSearch/config"
	_ "carSearch/docs"
	"carSearch/internal/adapter/database"
	"carSearch/internal/adapter/service"
	"carSearch/internal/exception"
	"carSearch/internal/handler/http"
	"carSearch/internal/repositories/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {

	// Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		Prefork:      false,
	})
	app.Use(recover.New())
	app.Use(cors.New())

	api := app.Group("/api")

	// Get DB
	db := database.NewDatabase(&config.Config{
		Database: config.Database{
			Host:     "postgres",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DBName:   "postgres",
		},
	})

	defer db.Close()

	//Repository
	carRepo := postgres.NewCarRepository(db.DB)
	manuRepo := postgres.NewManufacturerRepository(db.DB)
	// Service
	carService := service.NewCarService(carRepo, manuRepo)
	manufacturerService := service.NewManufactureService(manuRepo)
	// Routes
	carHandler := http.NewCarHandler(carService, manufacturerService)
	manufactureHandler := http.NewManufactureHandler(manufacturerService)

	http.NewHelpHandler().Route(api)
	carHandler.Route(api)
	manufactureHandler.Route(api)

	log.Fatal(app.Listen(":8000"))
}
