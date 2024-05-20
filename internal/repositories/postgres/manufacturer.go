package postgres

import (
	"carSearch/internal/model"
	"database/sql"
)

type manufacturerRepository struct {
	db *sql.DB
}

func NewManufacturerRepository(db *sql.DB) *manufacturerRepository {
	return &manufacturerRepository{db}
}

func (repository *manufacturerRepository) Create(manufacturer *model.Manufacturer) (int, error) {
	var id int
	err := repository.db.QueryRow("INSERT INTO manufacturer (name, hsn) VALUES ($1, $2) RETURNING id", manufacturer.Name, manufacturer.Hsn).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *manufacturerRepository) GetByHsn(hsn string) (*model.Manufacturer, error) {
	manufacturer := &model.Manufacturer{}
	err := repository.db.QueryRow("SELECT name, hsn, id FROM manufacturer WHERE hsn = $1", hsn).Scan(&manufacturer.Name, &manufacturer.Hsn, &manufacturer.ID)
	if err != nil {
		return nil, err
	}
	return manufacturer, nil
}
