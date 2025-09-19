package main

import "testing"

func TestAdd(t *testing.T) {
	result := 2 + 3
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	// t.Fatal() will stop the test immediately if it fails.
}
func TestSub(t *testing.T) {
	result := 4 - 1
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
func TestMul(t *testing.T) {
	result := 6 * 7
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}
func TestDiv(t *testing.T) {
	result := 8 / 2
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}
func TestMod(t *testing.T) {
	result := 9 % 4
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
