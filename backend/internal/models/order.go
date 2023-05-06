package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint `gorm:"primaryKey"`
	StoreId   uint
	Store     Store
	UserId    uint
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
	CartId    uint
	Cart      Cart
	Price     uint `gorm:"not null"`
	Status    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
