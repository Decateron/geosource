package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (location *Location) Scan(value interface{}) error {
	floatStrings := strings.Split(strings.Trim(string(value.([]byte)), "()"), ",")
	location.Latitude, _ = strconv.ParseFloat(floatStrings[0], 64)
	location.Longitude, _ = strconv.ParseFloat(floatStrings[1], 64)
	return nil
}

func (location Location) Value() (driver.Value, error) {
	return fmt.Sprintf("%f,%f", location.Latitude, location.Longitude), nil
}
