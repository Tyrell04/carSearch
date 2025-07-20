package domain

import (
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type Car struct {
	BaseModel
	HSN      string   `gorm:"size:255;not null"`
	TSN      string   `gorm:"size:255;not null"`
	Name     string   `gorm:"size:255;not null"`
	Producer Producer `gorm:"foreignKey:HSN;references:HSN"`
}

type Producer struct {
	BaseModel
	HSN  string `gorm:"size:255;not null;uniqueIndex"`
	Name string `gorm:"size:255;not null"`
}
