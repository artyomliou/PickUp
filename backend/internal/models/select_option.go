package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectOption struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	SelectQuestionId uint
	SelectQuestion   SelectQuestion
	Name             string `gorm:"not null"`
	Price            uint   `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
