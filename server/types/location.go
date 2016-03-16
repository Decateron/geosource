package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Validates the location, ensuring that its latitude and longitude fall within
// valid bounds.
func (location *Location) Validate() error {
	if location.Latitude > 90 || location.Latitude < -90 {
		return errors.New("Latitude must be between -90 and 90")
	}
	if location.Longitude > 180 || location.Longitude < -180 {
		return errors.New("Longitude must be between -180 and 180")
	}
	return nil
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
