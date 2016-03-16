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

// Reads a POINT from the database.
func (location *Location) Scan(value interface{}) error {
	floatStrings := strings.Split(strings.Trim(string(value.([]byte)), "()"), ",")
	location.Latitude, _ = strconv.ParseFloat(floatStrings[0], 64)
	location.Longitude, _ = strconv.ParseFloat(floatStrings[1], 64)
	return nil
}

// Converts the location into a POINT format which can be inserted into the
// database.
func (location Location) Value() (driver.Value, error) {
	return fmt.Sprintf("%f,%f", location.Latitude, location.Longitude), nil
}
