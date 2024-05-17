package main

import (
	_ "carSearch/docs"
	"carSearch/internal/handler/http"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	// Routes
	http.NewHelpHandler().Route(api)

	log.Fatal(app.Listen(":8000"))
}
