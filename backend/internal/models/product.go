package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	StoreId         uint
	Store           Store
	Categories      []*Category       `gorm:"many2many:category_product"`
	SelectQuestions []*SelectQuestion `gorm:"many2many:select_question_product"`
	CustomQuestions []*CustomQuestion `gorm:"many2many:custom_question_product"`
	Name            string            `gorm:"not null"`
	Price           uint              `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
