package service

import "carSearch/internal/model"

type ManufacturerRepository interface {
	Create(manufacturer *model.Manufacturer) (int, error)
	GetByHsn(hsn string) (*model.Manufacturer, error)
}

type manufactuerService struct {
	ManufacturerRepository
}

func NewManufactureService(manufacturerRepository ManufacturerRepository) *manufactuerService {
	return &manufactuerService{manufacturerRepository}
}

func (service *manufactuerService) GetByHsn(hsn string) (*model.Manufacturer, error) {
	manufacturer, err := service.ManufacturerRepository.GetByHsn(hsn)
	if err != nil {
		return nil, err
	}
	return manufacturer, nil
}
