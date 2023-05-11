package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Pic       string         `gorm:"not null"`
	Status    StoreStatus    `gorm:"type:uint;not null"`
	OpenedAt  string         `gorm:"size:5"`
	ClosedAt  string         `gorm:"size:5"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type StoreStatus string

const StoreStatusCreated string = "Created"
const StoreStatusOpened string = "Opened"
const StoreStatusClosed string = "Closed"
const StoreStatusTemporarilyClosed string = "Temporarily closed"

// https://gorm.io/docs/data_types.html
func (s *StoreStatus) Scan(value interface{}) error {
	intval, ok := value.(int64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to assert value was int:", value))
	}

	switch uint(intval) {
	case 1:
		*s = StoreStatus(StoreStatusCreated)
		return nil
	case 2:
		*s = StoreStatus(StoreStatusOpened)
		return nil
	case 3:
		*s = StoreStatus(StoreStatusClosed)
		return nil
	case 4:
		*s = StoreStatus(StoreStatusTemporarilyClosed)
		return nil
	default:
		return errors.New(fmt.Sprint("Unknown StoreStatus integer value:", intval))
	}
}

func (s StoreStatus) Value() (driver.Value, error) {
	switch string(s) {
	case StoreStatusCreated:
		return int64(1), nil
	case StoreStatusOpened:
		return int64(2), nil
	case StoreStatusClosed:
		return int64(3), nil
	case StoreStatusTemporarilyClosed:
		return int64(4), nil
	default:
		return nil, errors.New(fmt.Sprint("Unknown StoreStatus string value:", string(s)))
	}
}
