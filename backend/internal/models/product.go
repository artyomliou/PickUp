package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	StoreId   uint
	Store     Store
	ID        uint `gorm:"primaryKey"`
	Name      string
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
