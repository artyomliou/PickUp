package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement" form:"id"`
	CartId        uint           `json:"cartId"`
	Cart          Cart           `json:"-"`
	ProductId     uint           `json:"-"`
	Product       *Product       `json:"product"`
	Amount        uint           `json:"amount" form:"amount" binding:"required"`
	Total         int            `json:"total"`
	SelectAnswers SelectAnswers  `json:"selectAnswers" gorm:"serializer:json"`
	CustomAnswers CustomAnswers  `json:"customAnswers" gorm:"serializer:json"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}
