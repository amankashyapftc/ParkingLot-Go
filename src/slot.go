package src

import (
	"errors"

	"github.com/google/uuid"
)

type Token struct {
	Id string
}
type Slot struct {
	SlotNumber int
	ParkedCar  *Car
	Token      *Token
}

func NewSlot(slotNumber int, parkedCar *Car) *Slot {
	return &Slot{
		SlotNumber: slotNumber,
		ParkedCar:  parkedCar,
	}
}

func (s *Slot) Park(parkedCar *Car) (*Token, error) {
	if s.ParkedCar != nil {
		return nil, errors.New("slot is already occupied")
	}

	s.ParkedCar = parkedCar
	token := &Token{
		Id: uuid.New().String(),
	}
	s.Token = token
	return token, nil
}

func (s *Slot) IsCarPresent(targetCar *Car) bool {
	return s.ParkedCar != nil && s.ParkedCar == targetCar
}

func (s *Slot) isVerifiedToken(token *Token) bool {
	return s.Token == token
}
func (s *Slot) UnPark(token *Token) error {
	if s.ParkedCar == nil {
		return errors.New("slot is already unoccupied")
	}

	if token == nil || token.Id == "" {
		return errors.New("invalid token")
	}

	if s.isVerifiedToken(token) {
		s.ParkedCar = nil
	}
	return nil
}

func (s *Slot) IsOccupied() bool {
	return s.ParkedCar != nil
}
