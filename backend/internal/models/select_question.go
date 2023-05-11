package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectQuestion struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId    uint           `json:"-"`
	Store      Store          `json:"-"`
	Products   []*Product     `json:"-" gorm:"many2many:select_question_product"`
	Name       string         `json:"name" gorm:"not null"`
	IsRequired bool           `json:"isRequired" gorm:"not null;default=false"`
	AtMost     uint           `json:"atMost" gorm:"not null;default=0"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
