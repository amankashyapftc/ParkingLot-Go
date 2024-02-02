package src

import "testing"

func TestAbleToCreateANewAttendant(t *testing.T) {
	attendant, _ := NewAttendant()

	if attendant == nil {
		t.Errorf("Car should be created with parameters.")
	}
}

func TestAbleToParkCarUsingAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	car, _ := NewCar("HJ67YH6789", "Red")
	attendant.assign(parkingLot)

	attendant.park(car)

	if !parkingLot.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestAttendantShouldBeAbleParkInMultipleParkingLot(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot1, _ := NewParkingLot(1)
	parkingLot2, _ := NewParkingLot(1)
	car1, _ := NewCar("HJ67YH6789", "Red")
	car2, _ := NewCar("YU67GH67889", "Blue")
	attendant.assign(parkingLot1)
	attendant.assign(parkingLot2)

	attendant.park(car1)
	attendant.park(car2)

	if !parkingLot1.isFull() && !parkingLot2.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestCanNotParkCarIfAllParkingLotAreFull(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot1, _ := NewParkingLot(1)
	parkingLot2, _ := NewParkingLot(1)
	car1, _ := NewCar("HJ67YH6789", "Red")
	car2, _ := NewCar("YU67GH67889", "Blue")
	car3, _ := NewCar("GJ677Gh5678", "Green")
	attendant.assign(parkingLot1)
	attendant.assign(parkingLot2)

	attendant.park(car1)
	attendant.park(car2)
	_, err := attendant.park(car3)

	if err == nil {
		t.Errorf("Can not park car if all parking lot is full")
	}
}

func TestAbleToUnparkCarUsingAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	car, _ := NewCar("GJ677Gh5678", "Green")
	attendant.assign(parkingLot)

	token, _ := attendant.park(car)
	unparkedCar, _ := attendant.unPark(token)

	if car != unparkedCar {
		t.Errorf("Car should be unparked.")
	}
}

func TestCanNotUnparkCarUsingInValidToken(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	attendant.assign(parkingLot)

	_, err := attendant.unPark("random")

	if err == nil {
		t.Errorf("Can not un park car using wrong token.")
	}
}

func TestAttendantCanUnparkCarWhichIsParkByOtherAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	attendant1, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(2)
	attendant.assign(parkingLot)
	attendant1.assign(parkingLot)
	car, _ := NewCar("HJ67YH6789", "Red")

	token, _ := attendant.park(car)
	unparkedCar, _ := attendant1.unPark(token)

	if unparkedCar != car {
		t.Errorf("Attendant should be unpark car which parked by other attendant.")
	}
}

func TestParkCarNearestSlot(t *testing.T) {
	attendant, _ := NewAttendant(NEAREST)
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car, _ := NewCar("HJ67YH6789", "Red")
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)

	if !parkingLot.isFull() && parkingLot1.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestParkCarFarthestSlot(t *testing.T) {
	attendant, _ := NewAttendant(FARTHEST)
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car, _ := NewCar("HJ67YH6789", "Red")
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)

	if parkingLot.isFull() && !parkingLot1.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestParkCarsUsingDistributedStrategy(t *testing.T) {
	attendant, _ := NewAttendant(DISTRIBUTED)
	parkingLot, _ := NewParkingLot(2)
	parkingLot1, _ := NewParkingLot(2)
	car1, _ := NewCar("HJ67YH6789", "Red")
	car2, _ := NewCar("YU67GH67889", "Blue")
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car1)
	attendant.park(car2)

	if parkingLot.isFull() && parkingLot1.isFull() {
		t.Errorf("Cars should be park equaly in parking lots.")
	}
}
