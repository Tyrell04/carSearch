package http

import (
	"carSearch/internal/model"
	"github.com/gofiber/fiber/v2"
)

type CarService interface {
	Create(car *model.CarCreate) error
	ByHsnTsn(hsn, tsn string) (*model.Car, error)
}

type car struct {
	carService CarService
}

type Car struct {
	Name         string `json:"name"`
	Hsn          string `json:"hsn"`
	Tsn          string `json:"tsn"`
	Manufacturer string `json:"manufacturer"`
}

func NewCarHandler(carService CarService) *car {
	return &car{carService}
}

func (handler *car) Route(api fiber.Router) {
	api.Post("/car", handler.Create)
	api.Get("/car/:hsn/:tsn", handler.ByHsnTsn)
}

// Create
// @Summary
// @Description Create a new car
// @Tags car
// @Accept json
// @Produce json
// @Param Car body Car true "Car object"
// @Success 200
// @Router /api/car [post]
func (handler *car) Create(c *fiber.Ctx) error {
	car := new(Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := handler.carService.Create(&model.CarCreate{
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

// ByHsnTsn
// @Summary Get car by hsn and tsn
// @Description Get car by hsn and tsn
// @Tags car
// @Accept json
// @Produce json
// @Param hsn path string true "hsn"
// @Param tsn path string true "tsn"
// @Success 200
// @Router /api/car/{hsn}/{tsn} [get]
func (handler *car) ByHsnTsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	tsn := c.Params("tsn")

	car, err := handler.carService.ByHsnTsn(hsn, tsn)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(car)
}
