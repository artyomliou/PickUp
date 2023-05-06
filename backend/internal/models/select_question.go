package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectQuestion struct {
	ID         uint `gorm:"primaryKey"`
	StoreId    uint
	Store      Store
	Products   []*Product `gorm:"many2many:select_question_product"`
	Name       string     `gorm:"not null"`
	IsRequired bool       `gorm:"not null;default=false"`
	AtMost     uint       `gorm:"not null;default=0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
