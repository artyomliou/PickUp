package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Pic       string `gorm:"not null"`
	Status    uint   `gorm:"not null"`
	OpenedAt  string `gorm:"size:5"`
	ClosedAt  string `gorm:"size:5"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
