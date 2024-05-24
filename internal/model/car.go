package model

type Car struct {
	ID             int
	Name           string
	Tsn            string
	Hsn            string
	ManufacturerID int
}

type CarCreate struct {
	Name         string
	Tsn          string
	Hsn          string
	Manufacturer string
}
