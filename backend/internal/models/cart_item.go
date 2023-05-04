package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartId    uint `gorm:"primaryKey"`
	Cart      Cart
	ProductId uint `form:"productId" binding:"required" valid:"int"`
	Product   Product
	ID        uint `form:"id" binding:"optional" valid:"int"`
	Amount    uint `form:"amount" binding:"required" valid:"int"`
	Selects   SelectAnswers
	Customs   CustomAnswers
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItems []CartItem

func (data *CartItems) Value() (driver.Value, error) {
	b, err := json.Marshal(*data)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
func (data *CartItems) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return nil
}
func (data *CartItems) CalculateTotalPrice() uint {
	total := 0
	for _, item := range *data {
		total += int(item.Amount) * int(item.Product.Price)
	}
	return uint(total)
}
