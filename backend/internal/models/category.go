package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	StoreId   uint
	Store     Store
	Products  []*Product `gorm:"many2many:category_product"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
