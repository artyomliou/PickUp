package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement" form:"id"`
	CartId        uint           `json:"cartId"`
	Cart          Cart           `json:"-"`
	ProductId     uint           `json:"productId"`
	Product       Product        `json:"-"`
	Amount        uint           `json:"amount" form:"amount" binding:"required"`
	SelectAnswers SelectAnswers  `json:"selectAnswers" gorm:"type:json;not null;serializer:json"`
	CustomAnswers CustomAnswers  `json:"customAnswers" gorm:"type:json;not null;serializer:json"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}
