package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectOption struct {
	gorm.Model
	SelectQuestionId uint
	SelectQuestion   SelectQuestion
	ID               uint `gorm:"primaryKey"`
	Name             string
	Price            uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
