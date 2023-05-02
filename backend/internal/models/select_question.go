package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectQuestion struct {
	gorm.Model
	StoreId   uint
	Store     Store
	ID        uint `gorm:"primaryKey"`
	Name      string
	AtMost    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Products  []Product `gorm:"many2many:select_question_product"`
}
