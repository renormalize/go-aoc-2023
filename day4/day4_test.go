package day4

import "testing"

func TestSolveScratchCards(t *testing.T) {
	points, scratchCards, err := solveScratchCards("test_input.txt")
	if err != nil {
		t.Error("Solving for points and scratch cards failed with: ", err)
	}
	if points != 13 {
		t.Error("Expected 13, got ", points)
	}
	if scratchCards != 30 {
		t.Error("Expected 30, got ", scratchCards)
	}
}
