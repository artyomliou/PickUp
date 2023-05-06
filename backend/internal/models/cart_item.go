package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID        uint `gorm:"primaryKey;autoIncrement" form:"id"`
	CartId    uint
	Cart      Cart
	ProductId uint
	Product   Product
	Amount    uint `form:"amount" binding:"required"`
	Selects   SelectAnswers
	Customs   CustomAnswers
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
