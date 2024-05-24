package service

import "carSearch/internal/model"

type CarRepository interface {
	Create(car *model.Car) error
	ByHsnTsn(hsn, tsn string) (*model.Car, error)
}

type carService struct {
	CarRepository
	ManufacturerRepository
}

func NewCarService(repo CarRepository, manuRepo ManufacturerRepository) *carService {
	return &carService{repo, manuRepo}
}

func (service *carService) Create(car *model.CarCreate) error {
	manufacturer, err := service.ManufacturerRepository.GetByHsn(car.Hsn)
	if err != nil {
		id, err := service.ManufacturerRepository.Create(&model.Manufacturer{Hsn: car.Hsn, Name: car.Manufacturer})
		if err != nil {
			return err
		}
		manufacturer = &model.Manufacturer{ID: id, Hsn: car.Hsn, Name: car.Manufacturer}
	}
	return service.CarRepository.Create(&model.Car{Name: car.Name, Tsn: car.Tsn, ManufacturerID: manufacturer.ID})
}

func (service *carService) ByHsnTsn(hsn, tsn string) (*model.Car, error) {
	return service.CarRepository.ByHsnTsn(hsn, tsn)
}
