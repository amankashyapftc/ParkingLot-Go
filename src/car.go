package src

import (
	"errors"
	"regexp"
	"strings"
)

type Car struct {
	registrationNumber string
	color              string
}

func NewCar(registrationNumber, color string) (*Car, error) {
	if err := validateRegistrationNumber(registrationNumber, color); err != nil {
		return nil, err
	}
	return &Car{
		registrationNumber: registrationNumber,
		color:              color,
	}, nil
}

func validateRegistrationNumber(registrationNumber, color string) error {
	regex := "^[A-Za-z]{2}\\d{2}[A-Za-z]{2}\\d{4}$"
	if registrationNumber == "" || len(registrationNumber) < 10 || !regexp.MustCompile(regex).MatchString(registrationNumber) || color == "" {
		return errors.New("invalid registration number or color")
	}
	return nil
}

func (c *Car) Equals(other *Car) bool {
	return c.registrationNumber == other.registrationNumber && c.color == other.color
}

func (c *Car) HasColor(targetColor string) bool {
	return strings.EqualFold(c.color, targetColor)
}
