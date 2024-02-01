package src

import (
	"testing"
)

func TestAbleToCreateANewSlotWithNoCarParked(t *testing.T) {

	slot := NewSlot(1, nil)

	if slot == nil {
		t.Error("Expected a non-nil slot, but got nil")
	}
	if slot.SlotNumber != 1 {
		t.Errorf("Expected SlotNumber %d, got %d", 1, slot.SlotNumber)
	}
	if slot.ParkedCar != nil {
		t.Error("Expected ParkedCar to be nil, but it is not nil")
	}
}

func TestAbleToParkCarInBlankSlot(t *testing.T) {
	slot := NewSlot(1, nil)
	car, _ := NewCar("UP67GF4578", "Blue")
	slot.Park(car)
	expectedSlot := NewSlot(1, car)

	if slot == expectedSlot {
		t.Errorf("Expected slot %v, got %v", expectedSlot, slot)
	}
}

func TestAbleToGetTokenAfterParkACar(t *testing.T) {
	slot := NewSlot(1, nil)
	car, _ := NewCar("UP67TH3456", "Blue")

	token, _ := slot.Park(car)

	if token == nil {
		t.Errorf("Expected to get a token after parking a car, but got nil")
	}
}

func TestIfYouParkCarInAlreadyFilledSlotShouldThrowAnError(t *testing.T) {
	slot := NewSlot(1, nil)
	firstCar, _ := NewCar("LK34GH6754", "Blue")
	secondCar, _ := NewCar("UP30GT4532", "Red")

	_, err := slot.Park(firstCar)
	if err != nil {
		t.Errorf("Unexpected error while parking in an empty slot: %v", err)
	}

	_, err = slot.Park(secondCar)
	if err == nil {
		t.Errorf("Expected to get an error while parking in an already filled slot, but got none")
	}
}

func TestCarIsPresentInSlotOrNot(t *testing.T) {
	parkingSlot := NewSlot(1, nil)
	car, _ := NewCar("UP78DJ3534", "Blue")

	parkingSlot.Park(car)

	if !parkingSlot.IsCarPresent(car) {
		t.Error("Expected car to be present in the slot, but it's not.")
	}
}

func TestAbleToUnParkCarAfterVerifyingToken(t *testing.T) {
	parkingSlot := NewSlot(1, nil)
	car, _ := NewCar("AD23AS2345", "Blue")

	token, _ := parkingSlot.Park(car)
	parkingSlot.UnPark(token)

	if parkingSlot.IsOccupied() {
		t.Error("Expected slot to be unoccupied after unparking, but it's occupied.")
	}
}
