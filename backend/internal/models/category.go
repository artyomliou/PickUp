package models

import (
	"pick-up/backend/internal/db"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId   uint           `json:"-"`
	Store     Store          `json:"-"`
	Products  []*Product     `json:"products" gorm:"many2many:category_product"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func NewCategory(storeId uint, categoryName string) (*Category, error) {
	c := Category{}
	c.Name = categoryName
	c.StoreId = storeId
	if err := db.Conn().Create(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
