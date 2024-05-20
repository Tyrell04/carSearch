package postgres

import (
	"carSearch/internal/model"
	"database/sql"
)

type carRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *carRepository {
	return &carRepository{db}
}

func (repository *carRepository) Create(car *model.Car) error {
	_, err := repository.db.Exec("INSERT INTO cars (name, tsn, manufacturer_id) VALUES ($1, $2, $3)", car.Name, car.Tsn, car.ManufacturerID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *carRepository) GetByHsnTsn(hsn, tsn string) (*model.Car, error) {
	return nil, nil
}
