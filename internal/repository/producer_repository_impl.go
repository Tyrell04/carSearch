package repository

import (
	"carSearch/internal/domain"
	"context"
	"gorm.io/gorm"
)

type producerRepository struct {
	DB *gorm.DB
}

func NewProducerRepository(db *gorm.DB) domain.ProducerRepository {
	return &producerRepository{DB: db}
}

func (r *producerRepository) Save(ctx context.Context, producer *domain.Producer) error {
	return r.DB.WithContext(ctx).Save(producer).Error
}

func (r *producerRepository) FindByID(ctx context.Context, id uint) (*domain.Producer, error) {
	var producer domain.Producer
	if err := r.DB.WithContext(ctx).First(&producer, id).Error; err != nil {
		return nil, err
	}
	return &producer, nil
}

func (r *producerRepository) FindByHSN(ctx context.Context, hsn string) (*domain.Producer, error) {
	var producer domain.Producer
	if err := r.DB.WithContext(ctx).Where("hsn = ?", hsn).First(&producer).Error; err != nil {
		return nil, err
	}
	return &producer, nil
}

func (r *producerRepository) FirstOrCreate(ctx context.Context, producer *domain.Producer) error {
	return r.DB.WithContext(ctx).Where("hsn = ?", producer.HSN).FirstOrCreate(producer).Error
}
