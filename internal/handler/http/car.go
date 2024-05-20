package http

import (
	"carSearch/internal/model"
	"github.com/gofiber/fiber/v2"
)

type CarService interface {
	Create(car *model.CarCreate) error
	GetCarByHsnTsn(hsn, tsn string) (*model.Car, error)
}

type CarHandler struct {
	carService CarService
}

type Car struct {
	Name         string `json:"name"`
	Hsn          string `json:"hsn"`
	Tsn          string `json:"tsn"`
	Manufacturer string `json:"manufacturer"`
}

func NewCarHandler(carService CarService) *CarHandler {
	return &CarHandler{carService}
}

func (h *CarHandler) Route(api fiber.Router) {
	api.Post("/car", h.CreateCar)
	api.Get("/car/:hsn/:tsn", h.GetCarByHsnTsn)
}

// CreateCar
// @Summary
// @Description Analyze a text query
// @Tags analyze
// @Accept json
// @Produce json
// @Param Car body Car true "Car object"
// @Success 200
// @Router /api/car [post]
func (h *CarHandler) CreateCar(c *fiber.Ctx) error {
	car := new(Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.carService.Create(&model.CarCreate{
		Name:         car.Name,
		Hsn:          car.Hsn,
		Tsn:          car.Tsn,
		Manufacturer: car.Manufacturer,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Car created successfully"})
}

func (h *CarHandler) GetCarByHsnTsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	tsn := c.Params("tsn")

	car, err := h.carService.GetCarByHsnTsn(hsn, tsn)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(car)
}
