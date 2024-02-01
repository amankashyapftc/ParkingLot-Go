package slot

import "parking_lot/car"

type Slot struct {
	SlotNumber int
	ParkedCar  *car.Car
}

func NewSlot(slotNumber int, parkedCar *car.Car) *Slot {
	return &Slot{
		SlotNumber: slotNumber,
		ParkedCar:  parkedCar,
	}
}
