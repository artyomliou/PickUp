package models

import (
	"database/sql/driver"
	"encoding/json"
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

type StoreStatus uint

const StoreStatusCreated uint = 1
const StoreStatusOpened uint = 2
const StoreStatusClosed uint = 3
const StoreStatusTemporarilyClosed uint = 4

// https://gorm.io/docs/data_types.html
func (s *StoreStatus) Scan(value interface{}) error {
	intval, ok := value.(int64)
	if !ok {
		return errors.New(fmt.Sprint("Failed to assert value was int:", value))
	}

	*s = StoreStatus(uint(intval))
	return nil
}

func (s StoreStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// https://pkg.go.dev/encoding/json
func (s StoreStatus) MarshalJSON() ([]byte, error) {
	strval, err := s.ToHumanReadable(uint(s))
	if err != nil {
		return nil, err
	}
	return json.Marshal(strval)
}

func (s *StoreStatus) UnmarshalJSON(b []byte) error {
	var strval string
	if err := json.Unmarshal(b, &strval); err != nil {
		return err
	}

	intval, err := s.FromHumanReadable(strval)
	if err != nil {
		return err
	}
	*s = StoreStatus(intval)
	return nil
}

func (StoreStatus) ToHumanReadable(intval uint) (string, error) {
	switch intval {
	case StoreStatusCreated:
		return "Created", nil
	case StoreStatusOpened:
		return "Opened", nil
	case StoreStatusClosed:
		return "Closed", nil
	case StoreStatusTemporarilyClosed:
		return "Temporarily closed", nil
	}
	return "", errors.New(fmt.Sprint("Unknown StoreStatus integer value:", intval))
}

func (StoreStatus) FromHumanReadable(strval string) (uint, error) {
	switch strval {
	case "Created":
		return StoreStatusCreated, nil
	case "Opened":
		return StoreStatusOpened, nil
	case "Closed":
		return StoreStatusClosed, nil
	case "Temporarily closed":
		return StoreStatusTemporarilyClosed, nil
	}
	return 0, errors.New(fmt.Sprint("Unknown StoreStatus string value:", strval))
}
