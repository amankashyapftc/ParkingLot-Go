package car

import (
	"testing"
)

func TestAbleToCreateNewCar(t *testing.T) {
	car, err := NewCar("UP78DJ3587", "Green")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if car.registrationNumber != "UP78DJ3587" || car.color != "Green" {
		t.Errorf("Expected registration number %s and color %s, got %s and %s",
			"UP78DJ3587", "Green", car.registrationNumber, car.color)
	}
}

func TestCarRegistrationNumberHasACertainPattern(t *testing.T) {
	_, err := NewCar("AB12CD5643", "Silver")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}

func TestCarRegistrationNumberHasAtLeast10Char(t *testing.T) {
	_, err := NewCar("AB12CD345", "Yellow")
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	_, err = NewCar("AB12CD3456", "Silver")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}

func TestCarIfRegistrationNumberDoesNotFollowThePatternItShouldThrowError(t *testing.T) {
	_, err := NewCar("A1B2C3D4E5", "Black")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestCarRegistrationNumberCanNotBeEmpty(t *testing.T) {
	_, err := NewCar("", "Yellow")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestCarColorCanNotBeEmpty(t *testing.T) {
	_, err := NewCar("AB12CD3456", "")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
