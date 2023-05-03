package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

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
