package day2

import "testing"

func TestSolveCubeConundrum(t *testing.T) {
	sumIDs, sumPowers, err := solveCubeConundrum("test_input.txt")
	if err != nil {
		t.Error("Solving cube conundrum failed with: ", err)
	}
	if sumIDs != 8 {
		t.Error("Expected 8, got ", sumIDs)
	}
	if sumPowers != 2286 {
		t.Error("Expected 2286, got ", sumPowers)
	}
}
