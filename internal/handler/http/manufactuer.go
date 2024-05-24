package http

import (
	"carSearch/internal/exception"
	"carSearch/internal/models"
	"github.com/gofiber/fiber/v2"
)

type ManufactureService interface {
	GetByHsn(hsn string) (*models.Manufacturer, error)
}

type manufactureHandler struct {
	ManufactureService
}

func NewManufactureHandler(service ManufactureService) *manufactureHandler {
	return &manufactureHandler{service}
}

func (handler *manufactureHandler) Route(api fiber.Router) {
	api.Get("/car/:hsn", handler.GetByHsn)
}

// GetByHsn
// @Summary Get car by hsn
// @Description Get car by hsn
// @Tags manufacturer
// @Accept json
// @Produce json
// @Param hsn path string true "hsn"
// @Success 200
// @Router /api/car/{hsn} [get]
func (handler *manufactureHandler) GetByHsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")

	car, err := handler.ManufactureService.GetByHsn(hsn)
	if err != nil {
		exception.Panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(car)
}
