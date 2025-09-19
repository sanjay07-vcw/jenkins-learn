package main

import "testing"

func TestAdd(t *testing.T) {
	result := 2 + 3
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	// t.Fatal() will stop the test immediately if it fails.
}
