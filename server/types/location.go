package types

import (
	"errors"
	"log"
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
	log.Println(string(value.([]byte)))
	floatStrings := strings.Split(strings.Trim(string(value.([]byte)), "POINT()"), " ")
	location.Longitude, _ = strconv.ParseFloat(floatStrings[0], 64)
	location.Latitude, _ = strconv.ParseFloat(floatStrings[1], 64)
	return nil
}
