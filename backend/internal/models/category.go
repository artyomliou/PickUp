package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId   uint           `json:"-"`
	Store     Store          `json:"-"`
	Products  []*Product     `json:"products" gorm:"many2many:category_product"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
