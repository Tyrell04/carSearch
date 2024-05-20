package service

import "carSearch/internal/model"

type CarRepository interface {
	Create(car *model.Car) error
	GetByHsnTsn(hsn, tsn string) (*model.Car, error)
}

type ManufacturerRepository interface {
	Create(manufacturer *model.Manufacturer) (int, error)
	GetByHsn(hsn string) (*model.Manufacturer, error)
}

type CarService struct {
	carRepo  CarRepository
	manuRepo ManufacturerRepository
}

func NewCarService(repo CarRepository, manuRepo ManufacturerRepository) *CarService {
	return &CarService{repo, manuRepo}
}

func (service *CarService) Create(car *model.CarCreate) error {
	manufacturer, err := service.manuRepo.GetByHsn(car.Hsn)
	if err != nil {
		id, err := service.manuRepo.Create(&model.Manufacturer{Hsn: car.Hsn, Name: car.Manufacturer})
		if err != nil {
			return err
		}
		manufacturer = &model.Manufacturer{ID: id, Hsn: car.Hsn, Name: car.Manufacturer}
	}
	return service.carRepo.Create(&model.Car{Name: car.Name, Tsn: car.Tsn, ManufacturerID: manufacturer.ID})
}

func (service *CarService) GetCarByHsnTsn(hsn, tsn string) (*model.Car, error) {
	return service.carRepo.GetByHsnTsn(hsn, tsn)
}
