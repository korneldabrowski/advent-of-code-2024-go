package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {

	actualDistance := findOutDistanceToSanta("test_input.txt")

	expectedDistance := 11

	if actualDistance != expectedDistance {
		t.Errorf("Expected distance to be %v, but got %v", expectedDistance, actualDistance)
	}
}
