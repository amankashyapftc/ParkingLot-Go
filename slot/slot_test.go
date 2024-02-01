package slot

import (
	"testing"
)

func TestAbleToCreateANewSlotWithNoCarParked(t *testing.T) {
	slotNumber := 1

	slot := NewSlot(slotNumber, nil)

	if slot == nil {
		t.Error("Expected a non-nil slot, but got nil")
	}
	if slot.SlotNumber != slotNumber {
		t.Errorf("Expected SlotNumber %d, got %d", slotNumber, slot.SlotNumber)
	}
	if slot.ParkedCar != nil {
		t.Error("Expected ParkedCar to be nil, but it is not nil")
	}
}
