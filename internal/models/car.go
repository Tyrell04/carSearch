package models

type Car struct {
	ID               int    `json:"-"`
	Name             string `json:"name" example:"M 100"`
	Tsn              string `json:"tsn" example:"096"`
	Hsn              string `json:"hsn" example:"0001"`
	ManufacturerName string `json:"manufacturer_name" example:"ADLERWERKE"`
	ManufacturerID   int    `json:"-"`
}

type CarCreate struct {
	Name         string
	Tsn          string
	Hsn          string
	Manufacturer string
}
