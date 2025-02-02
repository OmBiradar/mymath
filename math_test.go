package mymath

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(3, 2)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
