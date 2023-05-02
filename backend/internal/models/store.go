package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Pic       string
	Status    uint
	OpenedAt  string
	ClosedAt  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
