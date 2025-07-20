package repository

import (
	"carSearch/internal/domain"
	"context"
	"gorm.io/gorm"
)

type carRepository struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) domain.CarRepository {
	return &carRepository{DB: db}
}

func (r *carRepository) Save(ctx context.Context, car *domain.Car) error {
	return r.DB.WithContext(ctx).Save(car).Error
}

func (r *carRepository) FindByID(ctx context.Context, id uint) (*domain.Car, error) {
	var car domain.Car
	if err := r.DB.WithContext(ctx).First(&car, id).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *carRepository) FindByHSN(ctx context.Context, hsn string) ([]*domain.Car, error) {
	var cars []*domain.Car
	if err := r.DB.WithContext(ctx).Preload("Producer").Where("hsn = ?", hsn).Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *carRepository) FindByHSNAndTSN(ctx context.Context, hsn string, tsn string) (*domain.Car, error) {
	var car domain.Car
	if err := r.DB.WithContext(ctx).Preload("Producer").Where("hsn = ? AND tsn = ?", hsn, tsn).First(&car).Error; err != nil {
		return nil, err
	}
	return &car, nil
}
