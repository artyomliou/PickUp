package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	StoreId   uint
	Store     Store
	UserId    uint
	User      User
	ID        uint `gorm:"primaryKey"`
	Items     string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
