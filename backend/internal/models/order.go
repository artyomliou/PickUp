package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	StoreId   uint           `json:"storeId"`
	Store     Store          `json:"store,omitempty"`
	UserId    uint           `json:"userId"`
	User      User           `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CartId    uint           `json:"cartId"`
	Cart      Cart           `json:"cart,omitempty"`
	Price     uint           `json:"price" gorm:"not null"`
	Status    OrderStatus    `json:"status" gorm:"type:uint;not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type OrderStatus string

const OrderStatusCreated string = "Created"
const OrderStatusPreparing string = "Preparing"
const OrderStatusReady string = "Ready"
const OrderStatusPicked string = "Picked"
const OrderStatusFinished string = "Finished"
const OrderStatusCancelled string = "Cancelled"

// https://gorm.io/docs/data_types.html
func (s *OrderStatus) Scan(value interface{}) error {
	intval, ok := value.(int64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to assert value was int:", value))
	}

	switch uint(intval) {
	case 1:
		*s = OrderStatus(OrderStatusCreated)
		return nil
	case 2:
		*s = OrderStatus(OrderStatusPreparing)
		return nil
	case 3:
		*s = OrderStatus(OrderStatusReady)
		return nil
	case 4:
		*s = OrderStatus(OrderStatusPicked)
		return nil
	case 5:
		*s = OrderStatus(OrderStatusFinished)
		return nil
	case 6:
		*s = OrderStatus(OrderStatusCancelled)
		return nil
	default:
		return errors.New(fmt.Sprint("Unknown OrderStatus integer value:", intval))
	}
}

func (s OrderStatus) Value() (driver.Value, error) {
	switch string(s) {
	case OrderStatusCreated:
		return int64(1), nil
	case OrderStatusPreparing:
		return int64(2), nil
	case OrderStatusReady:
		return int64(3), nil
	case OrderStatusPicked:
		return int64(4), nil
	case OrderStatusFinished:
		return int64(5), nil
	case OrderStatusCancelled:
		return int64(6), nil
	default:
		return nil, errors.New(fmt.Sprint("Unknown OrderStatus string value:", string(s)))
	}
}
