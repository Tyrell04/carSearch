package service

import "carSearch/internal/models"

type CarRepository interface {
	Create(car *models.Car) error
	ByHsnTsn(hsn, tsn string) (*models.Car, error)
}

type carService struct {
	CarRepository
	ManufacturerRepository
}

func NewCarService(repo CarRepository, manuRepo ManufacturerRepository) *carService {
	return &carService{repo, manuRepo}
}

func (service *carService) Create(car *models.CarCreate) error {
	manufacturer, err := service.ManufacturerRepository.GetByHsn(car.Hsn)
	if err != nil {
		id, err := service.ManufacturerRepository.Create(&models.Manufacturer{Hsn: car.Hsn, Name: car.Manufacturer})
		if err != nil {
			return err
		}
		manufacturer = &models.Manufacturer{ID: id, Hsn: car.Hsn, Name: car.Manufacturer}
	}
	return service.CarRepository.Create(&models.Car{Name: car.Name, Tsn: car.Tsn, ManufacturerID: manufacturer.ID})
}

func (service *carService) ByHsnTsn(hsn, tsn string) (*models.Car, error) {
	return service.CarRepository.ByHsnTsn(hsn, tsn)
}
