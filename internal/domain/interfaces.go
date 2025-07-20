package domain

import (
	"context"
)

type CarRepository interface {
	Save(ctx context.Context, car *Car) error
	FindByID(ctx context.Context, id uint) (*Car, error)
	FindByHSN(ctx context.Context, hsn string) ([]*Car, error)
	FindByHSNAndTSN(ctx context.Context, hsn string, tsn string) (*Car, error)
	// Add other car-specific repository methods as needed
}

type ProducerRepository interface {
	Save(ctx context.Context, producer *Producer) error
	FindByID(ctx context.Context, id uint) (*Producer, error)
	FindByHSN(ctx context.Context, hsn string) (*Producer, error)
	FirstOrCreate(ctx context.Context, producer *Producer) error
	// Add other producer-specific repository methods as needed
}
