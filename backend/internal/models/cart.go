package models

import (
	"pick-up/backend/internal/db"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId   uint           `json:"-"`
	Store     Store          `json:"-"`
	UserId    uint           `json:"-"`
	User      User           `json:"-"`
	Items     []*CartItem    `json:"items"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (cart *Cart) CalculateTotalPrice() uint {
	total := 0
	for _, item := range cart.Items {
		total += int(item.Amount) * int(item.Product.Price)
	}
	return uint(total)
}

func GetCart(userId uint, storeId uint) (*Cart, error) {
	cart := Cart{}
	tx := db.Conn().Where(Cart{
		UserId:  userId,
		StoreId: storeId,
	}).FirstOrCreate(&cart)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &cart, nil
	}
}

func GetCartItem(itemId uint) (*CartItem, error) {
	item := CartItem{}
	tx := db.Conn().Where(CartItem{
		ID: itemId,
	}).FirstOrCreate(&item)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &item, nil
	}
}
