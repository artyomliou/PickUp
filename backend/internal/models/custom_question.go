package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomQuestion struct {
	gorm.Model
	StoreId     uint
	Store       Store
	ID          uint `gorm:"primaryKey"`
	Name        string
	Placeholder string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Products    []Product `gorm:"many2many:custom_question_product"`
}
