package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	Price     uint        `gorm:"not null"`
	Status    OrderStatus `gorm:"type:uint;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type OrderStatus uint

const OrderStatusCreated uint = 1
const OrderStatusPreparing uint = 2
const OrderStatusReady uint = 3
const OrderStatusPicked uint = 4
const OrderStatusFinished uint = 5
const OrderStatusCancelled uint = 6

// https://gorm.io/docs/data_types.html
func (s *OrderStatus) Scan(value interface{}) error {
	intval, ok := value.(int64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to assert value was int:", value))
	}

	*s = OrderStatus(uint(intval))
	return nil
}

func (s OrderStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// https://pkg.go.dev/encoding/json
func (s OrderStatus) MarshalJSON() ([]byte, error) {
	strval, err := s.ToHumanReadable(uint(s))
	if err != nil {
		return nil, err
	}
	return json.Marshal(strval)
}

func (s *OrderStatus) UnmarshalJSON(b []byte) error {
	var strval string
	if err := json.Unmarshal(b, &strval); err != nil {
		return err
	}

	intval, err := s.FromHumanReadable(strval)
	if err != nil {
		return err
	}
	*s = OrderStatus(intval)
	return nil
}

func (OrderStatus) ToHumanReadable(intval uint) (string, error) {
	switch intval {
	case OrderStatusCreated:
		return "Created", nil
	case OrderStatusPreparing:
		return "Preparing", nil
	case OrderStatusReady:
		return "Ready", nil
	case OrderStatusPicked:
		return "Picked", nil
	case OrderStatusFinished:
		return "Finished", nil
	case OrderStatusCancelled:
		return "Cancelled", nil
	}
	return "", errors.New(fmt.Sprint("Unknown OrderStatus integer value:", intval))
}

func (OrderStatus) FromHumanReadable(strval string) (uint, error) {
	switch strval {
	case "Created":
		return OrderStatusCreated, nil
	case "Preparing":
		return OrderStatusPreparing, nil
	case "Ready":
		return OrderStatusReady, nil
	case "Picked":
		return OrderStatusPicked, nil
	case "Finished":
		return OrderStatusFinished, nil
	case "Cancelled":
		return OrderStatusCancelled, nil
	}
	return 0, errors.New(fmt.Sprint("Unknown OrderStatus string value:", strval))
}
