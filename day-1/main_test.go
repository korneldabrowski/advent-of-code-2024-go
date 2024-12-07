package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {

	actualDistance, err := findOutDistanceToSanta("test_input.txt")

	expectedDistance := 11

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actualDistance != expectedDistance {
		t.Errorf("Expected distance to be %v, but got %v", expectedDistance, actualDistance)
	}
}
