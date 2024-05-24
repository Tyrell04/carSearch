package postgres

import (
	"carSearch/internal/models"
	"database/sql"
)

type carRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *carRepository {
	return &carRepository{db}
}

func (repository *carRepository) Create(car *models.Car) error {
	_, err := repository.db.Exec("INSERT INTO cars (name, tsn, manufacturer_id) VALUES ($1, $2, $3)", car.Name, car.Tsn, car.ManufacturerID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *carRepository) ByHsnTsn(hsn, tsn string) (*models.Car, error) {
	car := &models.Car{}

	err := repository.db.QueryRow("SELECT id, hsn FROM manufacturer WHERE hsn = $1", hsn).Scan(&car.ManufacturerID, &car.Hsn)
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow("SELECT id, name, tsn FROM cars WHERE tsn = $1 AND manufacturer_id = $2", tsn, car.ManufacturerID).Scan(&car.ID, &car.Name, &car.Tsn)
	if err != nil {
		return nil, err
	}
	return car, nil
}
