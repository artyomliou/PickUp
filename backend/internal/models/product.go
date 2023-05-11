package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID              uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId         uint              `json:"-"`
	Store           Store             `json:"-"`
	Categories      []*Category       `json:"-" gorm:"many2many:category_product"`
	SelectQuestions []*SelectQuestion `json:"selectQuestions,omitempty" gorm:"many2many:select_question_product"`
	CustomQuestions []*CustomQuestion `json:"customQuestions,omitempty" gorm:"many2many:custom_question_product"`
	Name            string            `json:"name" gorm:"not null"`
	Price           uint              `json:"price" gorm:"not null"`
	CreatedAt       time.Time         `json:"-"`
	UpdatedAt       time.Time         `json:"-"`
	DeletedAt       gorm.DeletedAt    `json:"-"`
}
