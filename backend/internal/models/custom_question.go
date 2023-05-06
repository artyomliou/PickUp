package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomQuestion struct {
	ID          uint `gorm:"primaryKey"`
	StoreId     uint
	Store       Store
	Products    []*Product `gorm:"many2many:custom_question_product"`
	Name        string     `gorm:"not null"`
	Placeholder string     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
