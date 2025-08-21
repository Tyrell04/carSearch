package handler

import (
	"carSearch/internal/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type CarHandler struct {
	Service service.CarService
	APIKey  string
}

func NewCarHandler(s service.CarService, apiKey string) *CarHandler {
	return &CarHandler{Service: s, APIKey: apiKey}
}

func (h *CarHandler) ImportCars(c *fiber.Ctx) error {
	// Check API key
	providedAPIKey := c.Get("X-API-Key")
	if providedAPIKey != h.APIKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Handle multipart/form-data request
	file, err := c.FormFile("csv_file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get CSV file from form data",
			"error":   err.Error(),
		})
	}

	openFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to open CSV file",
			"error":   err.Error(),
		})
	}
	defer openFile.Close()

	// Import cars using the service
	if err := h.Service.ImportCarsFromCSV(openFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to import cars from CSV",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cars imported successfully",
	})
}

// SearchCars handles query parameter based search
// HSN is mandatory, TSN is optional
func (h *CarHandler) SearchCars(c *fiber.Ctx) error {
	hsn := c.Query("hsn")
	tsn := strings.ToUpper(c.Query("tsn"))

	if hsn == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "HSN query parameter is required",
		})
	}

	ctx := c.Context()

	// If TSN is provided, search for specific car
	if tsn != "" {
		car, err := h.Service.FindCarByHSNAndTSN(ctx, hsn, tsn)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Car not found",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"car": car,
		})
	}

	// If only HSN is provided, return producer information
	producer, err := h.Service.FindProducerByHSN(ctx, hsn)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Producer not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"producer": producer,
	})
}

// RegisterRoutes registers the API routes for the CarHandler.
func RegisterRoutes(app *fiber.App, handler *CarHandler) {
	app.Post("/import-cars", handler.ImportCars)
	app.Get("/cars/search", handler.SearchCars) // New query parameter endpoint
}
