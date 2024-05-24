package service

import "carSearch/internal/models"

type ManufacturerRepository interface {
	Create(manufacturer *models.Manufacturer) (int, error)
	GetByHsn(hsn string) (*models.Manufacturer, error)
}

type manufactuerService struct {
	ManufacturerRepository
}

func NewManufactureService(manufacturerRepository ManufacturerRepository) *manufactuerService {
	return &manufactuerService{manufacturerRepository}
}

func (service *manufactuerService) GetByHsn(hsn string) (*models.Manufacturer, error) {
	manufacturer, err := service.ManufacturerRepository.GetByHsn(hsn)
	if err != nil {
		return nil, err
	}
	return manufacturer, nil
}
