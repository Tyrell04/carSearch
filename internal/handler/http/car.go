package http

import (
	"carSearch/internal/exception"
	"carSearch/internal/models"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

type CarService interface {
	Create(car *models.CarCreate) error
	CreateFromCSV(file *multipart.FileHeader) error
	ByHsnTsn(hsn, tsn string) (*models.Car, error)
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
	api.Post("/car/csv", handler.CreateFromCSV)
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
		exception.Panic(err)
	}

	err := handler.carService.Create(&models.CarCreate{
		Name:         car.Name,
		Hsn:          car.Hsn,
		Tsn:          car.Tsn,
		Manufacturer: car.Manufacturer,
	})
	if err != nil {
		exception.Panic(err)
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
		exception.Panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(car)
}

// CreateFromCSV
// @Summary Create cars from CSV
// @Description Create cars from CSV
// @Tags car
// @Accept json
// @Produce json
// @Param file formData file true "CSV file"
// @Success 200
// @Router /api/car/csv [post]
func (handler *car) CreateFromCSV(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		exception.Panic(err)
	}

	err = handler.carService.CreateFromCSV(file)

	if err != nil {
		exception.Panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cars created successfully"})

}
