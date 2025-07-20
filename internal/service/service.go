package service

import (
	"carSearch/internal/domain"
	"context"
	"io"
)

type CarService interface {
	ImportCarsFromCSV(reader io.Reader) error
	FindCarsByHSN(ctx context.Context, hsn string) ([]*domain.Car, error)
	FindCarByHSNAndTSN(ctx context.Context, hsn string, tsn string) (*domain.Car, error)
	FindProducerByHSN(ctx context.Context, hsn string) (*domain.Producer, error)
}
