package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	StoreId   uint `gorm:"primaryKey"`
	Store     Store
	UserId    uint `gorm:"primaryKey"`
	User      User
	Items     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
