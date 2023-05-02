package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	StoreId   uint
	Store     Store
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Products  []Product `gorm:"many2many:category_product"`
}
