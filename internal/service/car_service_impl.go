package service

import (
	"carSearch/internal/domain"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type carService struct {
	carRepo      domain.CarRepository
	producerRepo domain.ProducerRepository
}

func NewCarService(carRepo domain.CarRepository, producerRepo domain.ProducerRepository) CarService {
	return &carService{carRepo: carRepo, producerRepo: producerRepo}
}

func (s *carService) ImportCarsFromCSV(reader io.Reader) error {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ',' // Assuming space as delimiter based on example

	ctx := context.Background()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		if len(record) < 4 {
			return fmt.Errorf("invalid CSV record: %v", record)
		}

		producerHSN := record[0]
		producerName := record[1]
		carTSN := record[2]
		carName := strings.Join(record[3:], ",")

		producer := domain.Producer{HSN: producerHSN, Name: producerName}
		// Check if producer exists, if not, create it
		if err := s.producerRepo.FirstOrCreate(ctx, &producer); err != nil {
			return fmt.Errorf("failed to create or find producer: %w", err)
		}

		car := domain.Car{HSN: producerHSN, TSN: carTSN, Name: carName}
		if err := s.carRepo.Save(ctx, &car); err != nil {
			return fmt.Errorf("failed to create car: %w", err)
		}
	}

	return nil
}

func (s *carService) FindCarsByHSN(ctx context.Context, hsn string) ([]*domain.Car, error) {
	cars, err := s.carRepo.FindByHSN(ctx, hsn)
	if err != nil {
		return nil, fmt.Errorf("failed to find cars by HSN: %w", err)
	}
	return cars, nil
}

func (s *carService) FindCarByHSNAndTSN(ctx context.Context, hsn string, tsn string) (*domain.Car, error) {
	car, err := s.carRepo.FindByHSNAndTSN(ctx, hsn, tsn)
	if err != nil {
		return nil, fmt.Errorf("failed to find car by HSN and TSN: %w", err)
	}
	return car, nil
}

func (s *carService) FindProducerByHSN(ctx context.Context, hsn string) (*domain.Producer, error) {
	producer, err := s.producerRepo.FindByHSN(ctx, hsn)
	if err != nil {
		return nil, fmt.Errorf("failed to find producer by HSN: %w", err)
	}
	return producer, nil
}
