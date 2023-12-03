package day1

import (
	"testing"
)

func TestSumCalibrationValuesDigits(t *testing.T) {
	calibrationDigits, err := sumCalibrationValuesDigits("test_digits_input.txt")
	if err != nil {
		t.Error("sumCalibrationValuesDigits failed with error: ", err)
	}
	if calibrationDigits != 142 {
		t.Error("Expected 142, got ", calibrationDigits)
	}
}

func TestSumCalibrationValuesDigitsWords(t *testing.T) {
	calibrationDigits, err := sumCalibrationValuesDigitsWords("test_words_input.txt")
	if err != nil {
		t.Error("sumCalibrationValuesDigitsWords failed with error: ", err)
	}
	if calibrationDigits != 281 {
		t.Error("Expected 281, got ", calibrationDigits)
	}
}
