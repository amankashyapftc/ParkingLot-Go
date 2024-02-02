package src

import (
	"errors"

	"github.com/google/uuid"
)

type Slot struct {
	SlotNumber int
	ParkedCar  *Car
	Token      string
}

func NewSlot() *Slot {
	return &Slot{}
}

func (s *Slot) park(car *Car) (string, error) {
	if s.IsOccupied() {
		return "", errors.New("Slot is occupied")
	}

	s.ParkedCar = car
	token := uuid.New().String()
	s.Token = string(token)
	return s.Token, nil
}

func (s *Slot) isVerifiedToken(token string) bool {
	return s.Token == token
}

func (s *Slot) unPark(token string) (*Car, error) {
	if !s.IsOccupied() {
		return nil, errors.New("Slot is empty")
	}
	if !s.isVerifiedToken(token) {
		return nil, errors.New("invalid token")
	}

	res := s.ParkedCar
	s.ParkedCar = nil
	s.Token = ""
	return res, nil
}

func (s *Slot) IsOccupied() bool {
	return s.ParkedCar != nil
}
