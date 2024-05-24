package main

import (
	"carSearch/config"
	_ "carSearch/docs"
	"carSearch/internal/adapter/database"
	"carSearch/internal/adapter/service"
	"carSearch/internal/handler/http"
	"carSearch/internal/repositories/postgres"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	// Fiber instance
	app := fiber.New()
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
	carHandler := http.NewCarHandler(carService)
	manufactureHandler := http.NewManufactureHandler(manufacturerService)

	http.NewHelpHandler().Route(api)
	carHandler.Route(api)
	manufactureHandler.Route(api)

	log.Fatal(app.Listen(":8000"))
}
