package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomQuestion struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId     uint           `json:"-"`
	Store       Store          `json:"-"`
	Products    []*Product     `json:"-" gorm:"many2many:custom_question_product"`
	Name        string         `json:"name" gorm:"not null"`
	Placeholder string         `json:"placeholder" gorm:"not null"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
