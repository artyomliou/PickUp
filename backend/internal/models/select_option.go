package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectOption struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	SelectQuestionId uint           `json:"selectQuestionId"`
	SelectQuestion   SelectQuestion `json:"-"`
	Name             string         `json:"name" gorm:"not null"`
	Price            uint           `json:"price" gorm:"not null"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}
