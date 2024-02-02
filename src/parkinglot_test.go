package src

import (
	"testing"
)

func TestAbleToCreateANewParkinglot(t *testing.T) {
	parkingLot, err := NewParkingLot(2)

	if parkingLot == nil && err != nil {
		t.Errorf("Not able create the parking lot.")
	}
}

func TestTryToCreateParkingLotWithNegativeNumberOfSlotsThrowsError(t *testing.T) {
	parkingLot, err := NewParkingLot(-2)

	if err == nil && parkingLot != nil {
		t.Errorf("Not able create the parking lot.")
	}
}

func TestParkingLotShouldBeEmptyWhenCreated(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)

	if parkingLot.isFull() {
		t.Errorf("Parking lot should be full.")
	}
}

func TestParkCarACarInParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car, _ := NewCar("UP78DH5678", "Green")

	parkingLot.park(car, NEAREST)

	if !parkingLot.isFull() {
		t.Errorf("Parking lot should be full.")
	}
}

func TestParkCarInFullParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car1, _ := NewCar("UP78DH5678", "Green")
	car2, _ := NewCar("UP78DH5645", "Blue")

	parkingLot.park(car1, NEAREST)
	_, err := parkingLot.park(car2, NEAREST)

	if err == nil {
		t.Errorf("Can not park car in full parking lot.")
	}
}

func TestUnParkCarUsingParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car, _ := NewCar("UP78DH5678", "Green")

	ticket, _ := parkingLot.park(car, NEAREST)
	unparkedCar, _ := parkingLot.unPark(ticket)

	if car != unparkedCar {
		t.Errorf("Unparked car should be same as parked car.")
	}
}

func TestUnParkCarFromWrongParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car, _ := NewCar("UP78DH5678", "Green")

	ticket, _ := parkingLot.park(car, NEAREST)
	_, err := parkingLot1.unPark(ticket)

	if err == nil {
		t.Errorf("Can not unpark car from wrong parking lot.")
	}
}

func TestUnParkCarUsingWrongTicket(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car, _ := NewCar("UP78DH5678", "Green")

	parkingLot.park(car, NEAREST)
	_, err := parkingLot.unPark("Abc")

	if err == nil {
		t.Errorf("Can not unpark car using wrong ticket.")
	}
}

func TestEmptySlotCount(t *testing.T) {
	parkingLot, _ := NewParkingLot(2)
	car, _ := NewCar("UP78DH5678", "Green")

	parkingLot.park(car, NEAREST)

	if parkingLot.emptySlotCount() != 1 {
		t.Errorf("empty slots count should be 1.")
	}
}
