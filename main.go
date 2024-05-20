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
	database := database.NewDatabase(&config.Config{
		Database: config.Database{
			Host:     "postgres",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DBName:   "postgres",
		},
	})

	defer database.Close()

	//Repository
	carRepo := postgres.NewCarRepository(database.DB)
	manuRepo := postgres.NewManufacturerRepository(database.DB)
	// Service
	carService := service.NewCarService(carRepo, manuRepo)
	// Routes
	carHandler := http.NewCarHandler(carService)

	http.NewHelpHandler().Route(api)
	carHandler.Route(api)

	log.Fatal(app.Listen(":8000"))
}
