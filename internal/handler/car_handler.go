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

func (h *CarHandler) FindByHSN(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	if hsn == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "HSN parameter is required",
		})
	}

	ctx := c.Context()
	cars, err := h.Service.FindCarsByHSN(ctx, hsn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find cars",
			"error":   err.Error(),
		})
	}

	if len(cars) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No cars found with the specified HSN",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"cars": cars,
	})
}

func (h *CarHandler) FindByHSNAndTSN(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	tsn := strings.ToUpper(c.Params("tsn"))

	if hsn == "" || tsn == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "HSN and TSN parameters are required",
		})
	}

	ctx := c.Context()
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

// SearchCars handles query parameter based search
// HSN is mandatory, TSN is optional
func (h *CarHandler) SearchCars(c *fiber.Ctx) error {
	hsn := c.Query("hsn")
	tsn := strings.ToUpper(c.Params("tsn"))

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
	app.Get("/cars/hsn/:hsn", handler.FindByHSN)
	app.Get("/cars/hsn/:hsn/tsn/:tsn", handler.FindByHSNAndTSN)
	app.Get("/cars/search", handler.SearchCars) // New query parameter endpoint
}
