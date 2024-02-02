package src

import (
	"testing"
)

func TestAbleToCreateANewSlot(t *testing.T) {

	slot := NewSlot()

	if slot == nil {
		t.Errorf("Car should be created with parameters.")
	}

}
func TestNewSlotMustBeEmptyWhenCreated(t *testing.T) {
	slot := NewSlot()

	if slot.IsOccupied() {
		t.Errorf("Slot should be free, when created.")
	}
}

func TestShouldBeParkCarInBlankSlot(t *testing.T) {
	car, _ := NewCar("UP78GH4567", "Blue")
	slot := NewSlot()

	ticket, err := slot.park(car)
	if !slot.isVerifiedToken(ticket) && err == nil {
		t.Errorf("Ticket is not valid.")
	}
}

func TestAbleToGetTokenAfterParkACar(t *testing.T) {
	slot := NewSlot()
	car, _ := NewCar("UP67TH3456", "Blue")

	token, _ := slot.park(car)

	if token == "" {
		t.Errorf("Expected to get a token after parking a car, but got nil")
	}
}

func TestIfYouParkCarInAlreadyFilledSlotShouldThrowAnError(t *testing.T) {
	slot := NewSlot()
	firstCar, _ := NewCar("LK34GH6754", "Blue")
	secondCar, _ := NewCar("UP30GT4532", "Red")

	_, err := slot.park(firstCar)
	if err != nil {
		t.Errorf("Unexpected error while parking in an empty slot: %v", err)
	}

	_, err = slot.park(secondCar)
	if err == nil {
		t.Errorf("Expected to get an error while parking in an already filled slot, but got none")
	}
}

func TestShouldBeUnparkCar(t *testing.T) {
	car, _ := NewCar("LK34GH6754", "Blue")
	slot := NewSlot()

	token, _ := slot.park(car)
	unparkedCar, err := slot.unPark(token)

	if car != unparkedCar && err == nil && !slot.isVerifiedToken(token) {
		t.Errorf("Invalid car unparked.")
	}
}

func TestUnparkCarWithEmptySlot(t *testing.T) {
	slot := NewSlot()

	car, err := slot.unPark("abc")
	if err == nil && car != nil {
		t.Errorf("Couldn't unpark empty slot.")
	}
}

func TestUnparkCarWithInvalidTicket(t *testing.T) {
	car, _ := NewCar("UP56GH5678", "Green")
	slot := NewSlot()

	slot.park(car)
	car, err := slot.unPark("random")

	if err == nil && car != nil {
		t.Errorf("Couldn't unpark car using invalid ticket.")
	}
}
