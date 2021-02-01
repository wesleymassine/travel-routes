package exceptions

import (
	"errors"
	"fmt"
	"strings"
)

// CheckInputAirport validate code airport
func CheckInputAirport(airport string) error {
	if len(airport) != 3 {
		return errors.New("Airport code must contain only 3 characters")
	}

	return nil
}

// EqualsAirport Check if airport is the same
func EqualsAirport(from, to string) error {
	if strings.Contains(from, to) {
		return errors.New("From airport and destination cannot be the same")
	}

	return nil
}

// CheckFlightPrice validate price
func CheckFlightPrice(price int) error {
	if price <= 0 {
		return errors.New("Flight price cannot be 0")
	}

	return nil
}

//CheckAirportCode check airport exist
func CheckAirportCode(code string) error {
	return errors.New(fmt.Sprintf("Airport %s reported does not exist", code))
}

//FlightNotFound check aiports not found
func FlightNotFound() error {
	return errors.New("Informed flight not found")
}
